/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	uname "github.com/crossplane/upjet/v2/pkg/types/name"
	tfjson "github.com/hashicorp/terraform-json"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/buzz-IT/provider-databricks/config/cluster"
	"github.com/buzz-IT/provider-databricks/config/namespaced"

	"github.com/pkg/errors"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/buzz-IT/provider-databricks"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func getProviderSchema(s string) (*tfschema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &tfschema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProvider(fwProvider, sdkProvider, generationProvider, "databricks.crossplane.io")
	if err != nil {
		return nil, err
	}
	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespaced returns provider configuration for namespace-scoped resources
func GetProviderNamespaced(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProvider(fwProvider, sdkProvider, generationProvider, "databricks.m.crossplane.io")
	if err != nil {
		return nil, err
	}
	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

func getProvider(fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool, rootGroup string) (*config.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		sdkProvider = p
	}

	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithRootGroup(rootGroup),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithDefaultResourceOptions(ResourceConfigurator()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithFeaturesPackage("internal/features"),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithSchemaTraversers(&config.SingletonListEmbedder{}),
	)

	for _, r := range pc.Resources {
		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = uname.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}
	}

	setV1Beta1(pc)
	return pc, nil
}

func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		l[i] = name + "$"
		i++
	}
	return l
}

func setV1Beta1(pc *config.Provider) {
	for name, r := range pc.Resources {
		r.Version = "v1beta1"
		r.PreviousVersions = nil
		r.SetCRDStorageVersion(r.Version)
		r.Conversions = nil
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
		pc.Resources[name] = r
	}
}

/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	clusterv1beta1 "github.com/buzz-IT/provider-databricks/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/buzz-IT/provider-databricks/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal databricks credentials as JSON"
	errTenantIDNotSet       = "tenant ID must be set in ProviderConfig when credential source is OIDCTokenFile or Upbound"
	errClientIDNotSet       = "client ID must be set in ProviderConfig when credential source is OIDCTokenFile or Upbound"

	keyHost                     = "host"
	keyEnvironment              = "environment"
	keyAzureWorkspaceResourceID = "azure_workspace_resource_id"
	keyAzureUseMsi              = "azure_use_msi"
	keyAzureClientID            = "azure_client_id"
	keyAzureClientSecret        = "azure_client_secret"
	keyAzureTenantID            = "azure_tenant_id"
	keyAzureEnvironment         = "azure_environment"
	keyClientID                 = "client_id"
	keyClientSecret             = "client_secret"
	keyAccountID                = "account_id"
	keyAuthType                 = "auth_type"
	keyAuthToken                = "token"
	keyDatabricksIDTokenFile    = "databricks_id_token_filepath"
	keyGoogleCredentials        = "google_credentials"
	keyGoogleServiceAccount     = "google_service_account"

	defaultOidcTokenFilePath   = "/var/run/secrets/azure/tokens/azure-identity-token"
	envAzureFederatedTokenFile = "AZURE_FEDERATED_TOKEN_FILE"
)

var (
	credentialsSourceUserAssignedManagedIdentity   xpv2.CredentialsSource = "UserAssignedManagedIdentity"
	credentialsSourceSystemAssignedManagedIdentity xpv2.CredentialsSource = "SystemAssignedManagedIdentity"
	credentialsSourceOIDCTokenFile                 xpv2.CredentialsSource = "OIDCTokenFile"
	credentialsSourceUpbound                       xpv2.CredentialsSource = "Upbound"

	upboundProviderIdentityTokenFile = "/var/run/secrets/upbound.io/provider/token"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}

		switch pcSpec.Credentials.Source { //nolint:exhaustive
		case credentialsSourceSystemAssignedManagedIdentity, credentialsSourceUserAssignedManagedIdentity:
			err = msiAuth(pcSpec, &ps)
		case credentialsSourceOIDCTokenFile:
			err = oidcAuth(pcSpec, &ps)
		case credentialsSourceUpbound:
			err = upboundAuth(pcSpec, &ps)
		default:
			err = defaultAuth(ctx, pcSpec, &ps, client)
		}

		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to prepare terraform.Setup")
		}

		return ps, errors.Wrap(configureNoForkDBXClient(ctx, &ps, *tfProvider), "failed to configure the no-fork Databricks client")
	}
}

func defaultAuth(ctx context.Context, pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup, client client.Client) error { // nolint: gocyclo
	data, err := xpresource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}
	data = []byte(strings.TrimSpace(string(data)))
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}

	// set credentials configuration
	if v, ok := creds[keyHost]; ok {
		ps.Configuration[keyHost] = v
	}
	if v, ok := creds[keyAzureWorkspaceResourceID]; ok {
		ps.Configuration[keyAzureWorkspaceResourceID] = v
	}
	if v, ok := creds[keyAzureUseMsi]; ok {
		ps.Configuration[keyAzureUseMsi] = v
	}
	if v, ok := creds[keyAzureClientID]; ok {
		ps.Configuration[keyAzureClientID] = v
	}
	if v, ok := creds[keyAzureClientSecret]; ok {
		ps.Configuration[keyAzureClientSecret] = v
	}
	if v, ok := creds[keyAzureTenantID]; ok {
		ps.Configuration[keyAzureTenantID] = v
	}
	if v, ok := creds[keyAuthType]; ok {
		ps.Configuration[keyAuthType] = v
	}
	if v, ok := creds[keyAuthToken]; ok {
		ps.Configuration[keyAuthToken] = v
	}
	if v, ok := creds[keyClientID]; ok {
		ps.Configuration[keyClientID] = v
	}
	if v, ok := creds[keyClientSecret]; ok {
		ps.Configuration[keyClientSecret] = v
	}
	if v, ok := creds[keyAccountID]; ok {
		ps.Configuration[keyAccountID] = v
	}
	if v, ok := creds[keyEnvironment]; ok {
		ps.Configuration[keyEnvironment] = v
	}
	if v, ok := creds[keyAzureEnvironment]; ok {
		ps.Configuration[keyAzureEnvironment] = v
	}
	if v, ok := creds[keyGoogleCredentials]; ok {
		ps.Configuration[keyGoogleCredentials] = v
	}
	if v, ok := creds[keyGoogleServiceAccount]; ok {
		ps.Configuration[keyGoogleServiceAccount] = v
	}

	return nil
}

func applyAzureCommon(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) {
	if pcSpec.Host != nil && len(*pcSpec.Host) > 0 {
		ps.Configuration[keyHost] = *pcSpec.Host
	}
	if pcSpec.AzureWorkspaceResourceID != nil && len(*pcSpec.AzureWorkspaceResourceID) > 0 {
		ps.Configuration[keyAzureWorkspaceResourceID] = *pcSpec.AzureWorkspaceResourceID
	}
	if pcSpec.ClientID != nil && len(*pcSpec.ClientID) > 0 {
		ps.Configuration[keyAzureClientID] = *pcSpec.ClientID
	}
	if pcSpec.TenantID != nil && len(*pcSpec.TenantID) > 0 {
		ps.Configuration[keyAzureTenantID] = *pcSpec.TenantID
	}
	if pcSpec.Environment != nil && len(*pcSpec.Environment) > 0 {
		ps.Configuration[keyAzureEnvironment] = *pcSpec.Environment
	}
}

func msiAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	ps.Configuration[keyAzureUseMsi] = true
	ps.Configuration[keyAuthType] = "azure-msi"
	applyAzureCommon(pcSpec, ps)
	return nil
}

func oidcTokenFilePath(pcSpec *namespacedv1beta1.ProviderConfigSpec) string {
	if pcSpec.OidcTokenFilePath != nil && len(*pcSpec.OidcTokenFilePath) > 0 {
		return *pcSpec.OidcTokenFilePath
	}
	if tokenFile := os.Getenv(envAzureFederatedTokenFile); tokenFile != "" {
		return tokenFile
	}
	return defaultOidcTokenFilePath
}

func oidcAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	if pcSpec.TenantID == nil || len(*pcSpec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pcSpec.ClientID == nil || len(*pcSpec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	ps.Configuration[keyAuthType] = "github-oidc-azure"
	ps.Configuration[keyDatabricksIDTokenFile] = oidcTokenFilePath(pcSpec)
	applyAzureCommon(pcSpec, ps)
	return nil
}

func upboundAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	if pcSpec.TenantID == nil || len(*pcSpec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pcSpec.ClientID == nil || len(*pcSpec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	ps.Configuration[keyAuthType] = "github-oidc-azure"
	ps.Configuration[keyDatabricksIDTokenFile] = upboundProviderIdentityTokenFile
	applyAzureCommon(pcSpec, ps)
	return nil
}

func configureNoForkDBXClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func enrichLocalSecretRefs(pc *namespacedv1beta1.ProviderConfig, mg xpresource.Managed) {
	if pc != nil && pc.Spec.Credentials.SecretRef != nil {
		pc.Spec.Credentials.SecretRef.Namespace = mg.GetNamespace()
	}
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg xpresource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case xpresource.LegacyManaged: //nolint:staticcheck // Cluster-scoped generated resources still use the legacy managed interface.
		return resolveProviderConfigLegacy(ctx, crClient, managed)
	case xpresource.ModernManaged:
		return resolveProviderConfigModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed")
	}
}

func legacyToModernProviderConfigSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	// TODO(erhan): this is hacky and potentially lossy, generate or manually implement
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfigLegacy(ctx context.Context, client client.Client, mg xpresource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // Cluster-scoped generated resources still use the legacy managed interface.
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := xpresource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return legacyToModernProviderConfigSpec(pc)
}

func resolveProviderConfigModern(ctx context.Context, crClient client.Client, mg xpresource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrapf(err, "referenced provider config kind %q is invalid for %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}
	pcObj, ok := pcRuntimeObj.(xpresource.ProviderConfig)
	if !ok {
		return nil, errors.Errorf("referenced provider config kind %q is not a provider config type %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		enrichLocalSecretRefs(pc, mg)
		pcSpec = pc.Spec
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		// TODO(erhan)
		return nil, errors.New("unknown")
	}
	t := xpresource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}

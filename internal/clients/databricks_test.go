package clients

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"

	namespacedv1beta1 "github.com/buzz-IT/provider-databricks/apis/namespaced/v1beta1"
)

func Test_oidcAuth_tokenFilePath(t *testing.T) {
	tenantID, clientID := "tenant", "client"
	explicitPath := "/explicit/path/azure-identity-token"
	envPath := "/var/run/secrets/azure/wi/token/azure-identity-token"

	cases := map[string]struct {
		oidcTokenFilePath *string
		envValue          string
		envSet            bool
		want              string
	}{
		"explicit_path_wins_over_env": {
			oidcTokenFilePath: &explicitPath,
			envValue:          envPath,
			envSet:            true,
			want:              explicitPath,
		},
		"env_used_when_no_explicit_path": {
			envValue: envPath,
			envSet:   true,
			want:     envPath,
		},
		"falls_back_to_default_when_env_unset": {
			want: defaultOidcTokenFilePath,
		},
		"falls_back_to_default_when_env_empty": {
			envValue: "",
			envSet:   true,
			want:     defaultOidcTokenFilePath,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.envSet {
				t.Setenv(envAzureFederatedTokenFile, tc.envValue)
			}
			pcSpec := &namespacedv1beta1.ProviderConfigSpec{
				TenantID:          &tenantID,
				ClientID:          &clientID,
				OidcTokenFilePath: tc.oidcTokenFilePath,
			}
			ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
			if err := oidcAuth(pcSpec, ps); err != nil {
				t.Fatalf("oidcAuth() returned unexpected error: %v", err)
			}
			got, _ := ps.Configuration[keyDatabricksIDTokenFile].(string)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("databricks_id_token_filepath mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_msiAuth(t *testing.T) {
	host, tenantID, clientID, workspace, environment := "https://adb.example.net", "tenant", "client", "/subscriptions/x/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/ws", "public"
	pcSpec := &namespacedv1beta1.ProviderConfigSpec{
		Host:                     &host,
		TenantID:                 &tenantID,
		ClientID:                 &clientID,
		AzureWorkspaceResourceID: &workspace,
		Environment:              &environment,
	}
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	msiAuth(pcSpec, ps)

	want := terraform.ProviderConfiguration{
		keyAzureUseMsi:              true,
		keyAuthType:                 "azure-msi",
		keyHost:                     host,
		keyAzureTenantID:            tenantID,
		keyAzureClientID:            clientID,
		keyAzureWorkspaceResourceID: workspace,
		keyAzureEnvironment:         environment,
	}
	if diff := cmp.Diff(want, ps.Configuration); diff != "" {
		t.Errorf("MSI configuration mismatch (-want +got):\n%s", diff)
	}
}

func Test_upboundAuth(t *testing.T) {
	tenantID, clientID, environment := "tenant", "client", "public"
	pcSpec := &namespacedv1beta1.ProviderConfigSpec{
		TenantID:    &tenantID,
		ClientID:    &clientID,
		Environment: &environment,
	}
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	if err := upboundAuth(pcSpec, ps); err != nil {
		t.Fatalf("upboundAuth() returned unexpected error: %v", err)
	}

	want := terraform.ProviderConfiguration{
		keyAuthType:              "github-oidc-azure",
		keyDatabricksIDTokenFile: upboundProviderIdentityTokenFile,
		keyAzureTenantID:         tenantID,
		keyAzureClientID:         clientID,
		keyAzureEnvironment:      environment,
	}
	if diff := cmp.Diff(want, ps.Configuration); diff != "" {
		t.Errorf("Upbound configuration mismatch (-want +got):\n%s", diff)
	}
}

func Test_sourceAuth_requiredFields(t *testing.T) {
	fields := []struct {
		name string
		make func(*namespacedv1beta1.ProviderConfigSpec, *terraform.Setup) error
		want string
	}{
		{name: "OIDC tenant ID", make: oidcAuth, want: errTenantIDNotSet},
		{name: "Upbound tenant ID", make: upboundAuth, want: errTenantIDNotSet},
	}
	for _, tc := range fields {
		t.Run(tc.name, func(t *testing.T) {
			pcSpec := &namespacedv1beta1.ProviderConfigSpec{}
			ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
			if err := tc.make(pcSpec, ps); err == nil || err.Error() != tc.want {
				t.Fatalf("expected error %q, got %v", tc.want, err)
			}
		})
	}
}

package xpprovider

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// sdkV2ResourceFallbacks keeps migrated resources on the SDK v2 implementation
// when Upjet still reconciles them with the SDK client (nested-set hashing).
var sdkV2ResourceFallbacks = []string{
	"databricks_library",
}

// GetProvider returns the Plugin Framework and SDK v2 Databricks providers for Upjet no-fork.
func GetProvider(_ context.Context) (fwprovider.Provider, *schema.Provider, error) {
	return pluginfw.GetDatabricksProviderPluginFramework(pluginfw.WithSdkV2ResourceFallbacks(sdkV2ResourceFallbacks)),
		sdkv2.DatabricksProvider(sdkv2.WithSdkV2ResourceFallbacks(sdkV2ResourceFallbacks)), nil
}

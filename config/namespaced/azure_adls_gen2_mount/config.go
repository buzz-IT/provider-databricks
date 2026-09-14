package azure_adls_gen2_mount

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_azure_adls_gen2_mount", func(r *config.Resource) {
		r.ShortGroup = "storage"
	})
}

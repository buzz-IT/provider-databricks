package grants

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_grants", func(r *config.Resource) {
		r.ShortGroup = "unity"
		// Kind is GrantMap because controller-gen would skip CRD generation:
		// the plural of Grant is grants, which collides with databricks_grants.
		r.Kind = "GrantMap"
	})
}

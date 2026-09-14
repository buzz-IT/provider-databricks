package warehouses_default_warehouse_override

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_warehouses_default_warehouse_override", func(r *config.Resource) {
		r.ShortGroup = "sql"
	})
}

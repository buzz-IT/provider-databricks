package ai_gateway_mcp_service

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_ai_gateway_mcp_service", func(r *config.Resource) {
		r.ShortGroup = "ai"
	})
}

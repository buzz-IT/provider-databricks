package materialized_features_feature_tag

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_materialized_features_feature_tag", func(r *config.Resource) {
		r.ShortGroup = "mlflow"
	})
}

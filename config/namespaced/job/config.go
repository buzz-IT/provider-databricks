package job

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_job", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "format")

		sqlEndpoint := config.Reference{TerraformName: "databricks_sql_endpoint"}
		for _, path := range []string{
			"notebook_task.warehouse_id",
			"dbt_task.warehouse_id",
			"sql_task.warehouse_id",
			"task.sql_task.warehouse_id",
			"dashboard_task.warehouse_id",
			"power_bi_task.warehouse_id",
			"task.notebook_task.warehouse_id",
			"task.dbt_task.warehouse_id",
			"task.dashboard_task.warehouse_id",
			"task.power_bi_task.warehouse_id",
			"for_each_task.task.notebook_task.warehouse_id",
			"for_each_task.task.dbt_task.warehouse_id",
			"for_each_task.task.sql_task.warehouse_id",
			"for_each_task.task.dashboard_task.warehouse_id",
			"for_each_task.task.power_bi_task.warehouse_id",
		} {
			r.References[path] = sqlEndpoint
		}
	})
}

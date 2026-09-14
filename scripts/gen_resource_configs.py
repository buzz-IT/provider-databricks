#!/usr/bin/env python3
"""Generate Upjet resource configs, include lists, and provider.go registries."""

from __future__ import annotations

from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
MODULE = "github.com/buzz-IT/provider-databricks"

# SDK v2 resources from terraform-provider-databricks v1.132.0
# (WorkspaceResources + AccountResources + DualResources + AllSettingsResources).
SDK_RESOURCES: dict[str, str] = {
    "access_control_rule_set": "security",
    "aibi_dashboard_embedding_access_policy_setting": "settings",
    "aibi_dashboard_embedding_approved_domains_setting": "settings",
    "alert": "sql",
    "artifact_allowlist": "unity",
    "automatic_cluster_update_workspace_setting": "settings",
    "aws_s3_mount": "storage",
    "azure_adls_gen1_mount": "storage",
    "azure_adls_gen2_mount": "storage",
    "azure_blob_mount": "storage",
    "budget": "billing",
    "catalog": "unity",
    "catalog_workspace_binding": "unity",
    "cluster": "compute",
    "cluster_policy": "compute",
    "compliance_security_profile_workspace_setting": "settings",
    "connection": "unity",
    "credential": "unity",
    "custom_app_integration": "apps",
    "dashboard": "sql",
    "dbfs_file": "storage",
    "default_namespace_setting": "settings",
    "directory": "workspace",
    "disable_legacy_access_setting": "settings",
    "disable_legacy_dbfs_setting": "settings",
    "disable_legacy_features_setting": "settings",
    "enhanced_security_monitoring_workspace_setting": "settings",
    "entitlements": "security",
    "external_location": "unity",
    "file": "storage",
    "git_credential": "workspace",
    "global_init_script": "workspace",
    "grant": "unity",
    "grants": "unity",
    "group": "security",
    "group_instance_profile": "security",
    "group_member": "security",
    "group_role": "security",
    "instance_pool": "compute",
    "instance_profile": "security",
    "ip_access_list": "security",
    "job": "compute",
    "lakehouse_monitor": "unity",
    "library": "compute",
    "metastore": "unity",
    "metastore_assignment": "unity",
    "metastore_data_access": "unity",
    "mlflow_experiment": "mlflow",
    "mlflow_model": "mlflow",
    "mlflow_webhook": "mlflow",
    "model_serving": "serving",
    "model_serving_provisioned_throughput": "serving",
    "mount": "storage",
    "mws_credentials": "deployment",
    "mws_customer_managed_keys": "deployment",
    "mws_log_delivery": "log",
    "mws_ncc_binding": "deployment",
    "mws_ncc_private_endpoint_rule": "deployment",
    "mws_network_connectivity_config": "deployment",
    "mws_networks": "deployment",
    "mws_permission_assignment": "security",
    "mws_private_access_settings": "deployment",
    "mws_storage_configurations": "deployment",
    "mws_vpc_endpoint": "deployment",
    "mws_workspaces": "deployment",
    "notebook": "workspace",
    "notification_destination": "workspace",
    "obo_token": "security",
    "online_table": "unity",
    "permission_assignment": "security",
    "permissions": "security",
    "pipeline": "compute",
    "provider": "sharing",
    "query": "sql",
    "recipient": "sharing",
    "registered_model": "mlflow",
    "repo": "workspace",
    "restrict_workspace_admins_setting": "settings",
    "schema": "unity",
    "secret": "security",
    "secret_acl": "security",
    "secret_scope": "security",
    "service_principal": "security",
    "service_principal_role": "security",
    "service_principal_secret": "security",
    "sql_alert": "sql",
    "sql_dashboard": "sql",
    "sql_endpoint": "sql",
    "sql_global_config": "sql",
    "sql_permissions": "sql",
    "sql_query": "sql",
    "sql_table": "unity",
    "sql_visualization": "sql",
    "sql_widget": "sql",
    "storage_credential": "unity",
    "system_schema": "unity",
    "table": "unity",
    "token": "security",
    "user": "security",
    "user_instance_profile": "security",
    "user_role": "security",
    "vector_search_endpoint": "serving",
    "vector_search_index": "serving",
    "volume": "unity",
    "workspace_binding": "unity",
    "workspace_conf": "workspace",
    "workspace_file": "workspace",
}

# Plugin Framework resources served by default in v1.132.0.
# Migrated SDK resources (quality_monitor, share) are FW-only at runtime.
# library stays SDK-only because of nested-type hashing in the SDK schema.
FW_RESOURCES: dict[str, str] = {
    "account_federation_policy": "oauth",
    "account_iam_direct_group_member_v2": "security",
    "account_iam_group_v2": "security",
    "account_iam_service_principal_v2": "security",
    "account_iam_user_v2": "security",
    "account_iam_workspace_assignment_v2": "security",
    "account_network_policy": "settings",
    "account_setting_user_preference_v2": "settings",
    "account_setting_v2": "settings",
    "ai_gateway_mcp_service": "ai",
    "ai_gateway_model_provider_service": "ai",
    "ai_gateway_model_service": "ai",
    "ai_search_endpoint": "ai",
    "ai_search_index": "ai",
    "alert_v2": "sql",
    "app": "apps",
    "app_space": "apps",
    "apps_settings_custom_template": "apps",
    "budget_policy": "billing",
    "data_classification_catalog_config": "governance",
    "data_quality_monitor": "unity",
    "data_quality_refresh": "unity",
    "database_database_catalog": "databases",
    "database_instance": "databases",
    "database_synced_database_table": "databases",
    "disaster_recovery_failover_group": "dr",
    "disaster_recovery_stable_url": "dr",
    "domain": "unity",
    "endpoint": "serving",
    "entity_tag_assignment": "unity",
    "environments_default_workspace_base_environment": "envs",
    "environments_workspace_base_environment": "envs",
    "external_metadata": "unity",
    "feature_engineering_feature": "mlflow",
    "feature_engineering_kafka_config": "mlflow",
    "feature_engineering_materialized_feature": "mlflow",
    "knowledge_assistant": "ai",
    "knowledge_assistant_knowledge_source": "ai",
    "materialized_features_feature_tag": "mlflow",
    "online_store": "serving",
    "policy_info": "governance",
    "postgres_branch": "postgres",
    "postgres_catalog": "postgres",
    "postgres_cdf_config": "postgres",
    "postgres_data_api": "postgres",
    "postgres_database": "postgres",
    "postgres_endpoint": "postgres",
    "postgres_project": "postgres",
    "postgres_role": "postgres",
    "postgres_snapshot_schedule": "postgres",
    "postgres_synced_table": "postgres",
    "quality_monitor": "unity",
    "quality_monitor_v2": "unity",
    "rfa_access_request_destinations": "unity",
    "sandbox": "compute",
    "secret_uc": "unity",
    "service_principal_federation_policy": "oauth",
    "share": "sharing",
    "supervisor_agent": "ai",
    "supervisor_agent_tool": "ai",
    "tag_policy": "tags",
    "warehouses_default_warehouse_override": "sql",
    "workspace_entity_tag_assignment": "tags",
    "workspace_iam_direct_group_member_v2": "security",
    "workspace_iam_group_v2": "security",
    "workspace_iam_service_principal_v2": "security",
    "workspace_iam_user_v2": "security",
    "workspace_iam_workspace_assignment_v2": "security",
    "workspace_iam_workspace_identity_detail_v2": "security",
    "workspace_network_option": "settings",
    "workspace_setting_v2": "settings",
}

CONFIG_GO = '''package {pkg}

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {{
	p.AddResourceConfigurator("databricks_{pkg}", func(r *config.Resource) {{
		r.ShortGroup = "{group}"
	}})
}}
'''

PROVIDER_GO = '''// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package {scope}

import (
{imports}
)

func init() {{
{adds}
}}
'''


def tf_name(base: str) -> str:
    return f"databricks_{base}"


def ident_line(name: str) -> str:
    return f'\t"{name}": config.IdentifierFromProvider,'


def write_external_name() -> None:
    sdk_lines = "\n".join(ident_line(tf_name(n)) for n in sorted(SDK_RESOURCES))
    fw_lines = "\n".join(ident_line(tf_name(n)) for n in sorted(FW_RESOURCES))

    content = f'''/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{{
{sdk_lines}
}}

var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{{
{fw_lines}
}}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{{}}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {{
	return func(r *config.Resource) {{
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {{
			r.ExternalName = e
		}}
		if e, ok := TerraformPluginFrameworkExternalNameConfigs[r.Name]; ok {{
			r.ExternalName = e
		}}
	}}
}}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {{
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {{
		l[i] = name + "$"
		i++
	}}
	return l
}}

// ResourceConfigurator applies all external name configs listed in the SDK, FW
// and CLI tables.
func ResourceConfigurator() config.ResourceOption {{
	return func(r *config.Resource) {{
		e, configured := TerraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {{
			e, configured = TerraformPluginFrameworkExternalNameConfigs[r.Name]
		}}
		if !configured {{
			e, configured = CLIReconciledExternalNameConfigs[r.Name]
		}}
		if !configured {{
			return
		}}
		r.ExternalName = e
	}}
}}
'''
    (ROOT / "config" / "external_name.go").write_text(content)


def ensure_config(scope: str, base: str, group: str) -> None:
    path = ROOT / "config" / scope / base / "config.go"
    if path.exists():
        return
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(CONFIG_GO.format(pkg=base, group=group))


def write_provider_go(scope: str) -> None:
    cfg_dir = ROOT / "config" / scope
    packages = sorted(
        p.name
        for p in cfg_dir.iterdir()
        if p.is_dir() and (p / "config.go").exists() and p.name not in {"common"}
    )
    imports = "\n".join(
        f'\t"{MODULE}/config/{scope}/{pkg}"' for pkg in packages
    )
    adds = "\n".join(f"\tProviderConfiguration.AddConfig({pkg}.Configure)" for pkg in packages)
    (cfg_dir / "provider.go").write_text(
        PROVIDER_GO.format(scope=scope, imports=imports, adds=adds)
    )


def main() -> None:
    overlap = set(SDK_RESOURCES) & set(FW_RESOURCES)
    if overlap:
        raise SystemExit(f"resource listed as both SDK and FW: {sorted(overlap)}")

    all_resources = dict(SDK_RESOURCES)
    all_resources.update(FW_RESOURCES)
    for base, group in all_resources.items():
        ensure_config("cluster", base, group)
        ensure_config("namespaced", base, group)

    write_provider_go("cluster")
    write_provider_go("namespaced")
    write_external_name()
    print(f"SDK resources: {len(SDK_RESOURCES)}")
    print(f"FW resources:  {len(FW_RESOURCES)}")
    print(f"total:         {len(all_resources)}")


if __name__ == "__main__":
    main()

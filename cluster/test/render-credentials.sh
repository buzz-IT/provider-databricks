#!/usr/bin/env bash
# Build Databricks provider credentials JSON from env/secrets.
# Prints JSON to stdout. Empty values are omitted.
set -euo pipefail

python3 - <<'PY'
import json
import os

keys = [
    "host",
    "token",
    "account_id",
    "client_id",
    "client_secret",
    "auth_type",
    "azure_workspace_resource_id",
    "azure_client_id",
    "azure_client_secret",
    "azure_tenant_id",
    "azure_environment",
    "azure_use_msi",
    "google_credentials",
    "google_service_account",
    "environment",
]

env_map = {
    "host": "DATABRICKS_HOST",
    "token": "DATABRICKS_TOKEN",
    "account_id": "DATABRICKS_ACCOUNT_ID",
    "client_id": "DATABRICKS_CLIENT_ID",
    "client_secret": "DATABRICKS_CLIENT_SECRET",
    "auth_type": "DATABRICKS_AUTH_TYPE",
    "azure_workspace_resource_id": "DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID",
    "azure_client_id": "DATABRICKS_AZURE_CLIENT_ID",
    "azure_client_secret": "DATABRICKS_AZURE_CLIENT_SECRET",
    "azure_tenant_id": "DATABRICKS_AZURE_TENANT_ID",
    "azure_environment": "DATABRICKS_AZURE_ENVIRONMENT",
    "azure_use_msi": "DATABRICKS_AZURE_USE_MSI",
    "google_credentials": "DATABRICKS_GOOGLE_CREDENTIALS",
    "google_service_account": "DATABRICKS_GOOGLE_SERVICE_ACCOUNT",
    "environment": "DATABRICKS_ENVIRONMENT",
}

raw = os.environ.get("UPTEST_CLOUD_CREDENTIALS", "").strip()
if raw:
    json.loads(raw)  # validate
    print(raw)
    raise SystemExit(0)

creds = {}
for key in keys:
    val = os.environ.get(env_map[key], "").strip()
    if val:
        creds[key] = val

if "host" not in creds:
    raise SystemExit("DATABRICKS_HOST (or UPTEST_CLOUD_CREDENTIALS) is required")
if "token" not in creds and not (("client_id" in creds and "client_secret" in creds) or "google_credentials" in creds or "azure_client_id" in creds):
    raise SystemExit("set DATABRICKS_TOKEN, or M2M/Azure/GCP credentials, or UPTEST_CLOUD_CREDENTIALS")

print(json.dumps(creds, separators=(",", ":")))
PY

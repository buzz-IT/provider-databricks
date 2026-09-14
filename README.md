# Provider Databricks

Repository: [github.com/buzz-IT/provider-databricks](https://github.com/buzz-IT/provider-databricks)

`provider-databricks` is a [Crossplane](https://crossplane.io/) provider built
with [Upjet](https://github.com/crossplane/upjet). It exposes XRM-conformant
managed resources for the
[Databricks Terraform provider](https://registry.terraform.io/providers/databricks/databricks/latest/docs).

The provider is generated against **terraform-provider-databricks v1.132.0**
using **Upjet v2** (cluster-scoped and namespaced MRs, Plugin Framework + SDK v2
no-fork runtime, SafeStart, management policies).

## Install

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-databricks
spec:
  package: ghcr.io/buzz-it/provider-databricks:v0.1.0
```

## Configure

Create a secret with Databricks credentials, then a `ProviderConfig`.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: databricks-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "host": "https://adb-xxxxxxxx.azuredatabricks.net",
      "token": "dapiXXXXXXXX"
    }
---
apiVersion: databricks.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: databricks-creds
      key: credentials
```

Supported credential JSON keys match the Terraform provider, including
`host`, `token`, `account_id`, `client_id`, `client_secret`,
`azure_workspace_resource_id`, `azure_client_id`, `azure_client_secret`,
`azure_tenant_id`, `google_credentials`, and `google_service_account`.

Cluster-scoped managed resources use `ProviderConfig`. Namespaced managed
resources (`*.databricks.m.crossplane.io`) can reference a namespaced
`ProviderConfig` or a cluster-scoped `ClusterProviderConfig`.

## Develop

```console
make generate
make run
make build
```

`make generate` clones `terraform-provider-databricks` v1.132.0 into
`hack/terraform-provider-databricks`, injects the exported `xpprovider`
package required for Upjet's no-fork architecture, scrapes resource docs, and
generates CRDs, controllers, and examples.

## End-to-end tests

`make e2e` starts Kind, installs Crossplane, builds this provider, and runs
[uptest](https://github.com/crossplane/uptest) against real Databricks APIs.

Default suite (cheap, workspace-level, no clusters):

- `examples/e2e/cluster/secretscope.yaml`
- `examples/e2e/cluster/directory.yaml`
- `examples/e2e/namespaced/secretscope.yaml`
- `examples/e2e/namespaced/directory.yaml`

Optional suites (set `UPTEST_EXAMPLE_LIST`):

| Suite | Paths |
| --- | --- |
| Groups | `examples/e2e/cluster/group.yaml,examples/e2e/namespaced/group.yaml` |
| Unity Catalog | `examples/e2e/cluster/catalog.yaml,examples/e2e/namespaced/catalog.yaml` |
| Compute | `examples/e2e/cluster/cluster.yaml,examples/e2e/namespaced/cluster.yaml` |

### Credentials

Set **either** a single JSON blob **or** individual variables. Individual
values are assembled into the provider secret automatically.

| Name | Secret or var | Required | Purpose |
| --- | --- | --- | --- |
| `DATABRICKS_HOST` | secret or var | yes* | Workspace or account URL |
| `DATABRICKS_TOKEN` | secret | yes* | PAT |
| `DATABRICKS_ACCOUNT_ID` | secret or var | no | Account-level APIs |
| `DATABRICKS_CLIENT_ID` | secret or var | no | M2M OAuth |
| `DATABRICKS_CLIENT_SECRET` | secret | no | M2M OAuth |
| `DATABRICKS_AUTH_TYPE` | secret or var | no | Terraform `auth_type` |
| `DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID` | secret or var | no | Azure Databricks |
| `DATABRICKS_AZURE_CLIENT_ID` | secret or var | no | Azure SP |
| `DATABRICKS_AZURE_CLIENT_SECRET` | secret | no | Azure SP |
| `DATABRICKS_AZURE_TENANT_ID` | secret or var | no | Azure SP |
| `DATABRICKS_AZURE_ENVIRONMENT` | var | no | Azure cloud name |
| `DATABRICKS_AZURE_USE_MSI` | var | no | `"true"` to use MSI |
| `DATABRICKS_GOOGLE_CREDENTIALS` | secret | no | GCP JSON key |
| `DATABRICKS_GOOGLE_SERVICE_ACCOUNT` | secret or var | no | GCP SA email |
| `UPTEST_CLOUD_CREDENTIALS` | secret | no | Raw provider JSON; overrides the fields above |

\*Required unless you use M2M/Azure/GCP instead of a PAT, or set
`UPTEST_CLOUD_CREDENTIALS`.

### Resource parameters (Actions variables)

| Name | Default |
| --- | --- |
| `DATABRICKS_NODE_TYPE_ID` | `Standard_DS3_v2` |
| `DATABRICKS_SPARK_VERSION` | `15.4.x-scala2.12` |
| `DATABRICKS_CATALOG_NAME` | `uptest_e2e` |
| `DATABRICKS_SCHEMA_NAME` | `uptest` |
| `DATABRICKS_DIRECTORY_PATH` | `/uptest-e2e` |
| `DATABRICKS_GROUP_DISPLAY_NAME` | `uptest-e2e-group` |
| `UPTEST_EXAMPLE_LIST` | smoke suite above |
| `UPTEST_SKIP_DELETE` | unset (resources are deleted) |

GitHub Actions: repository **Settings → Secrets and variables → Actions**.
The `E2E` workflow reads secrets first, then variables. Manual runs use
**Actions → E2E → Run workflow**. PRs and the Monday schedule skip (or, for
manual runs, fail) when credentials are missing.

Local:

```console
export DATABRICKS_HOST="https://adb-xxxxxxxx.azuredatabricks.net"
export DATABRICKS_TOKEN="dapi..."
make e2e
```

## License

Apache-2.0

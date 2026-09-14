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

## License

Apache-2.0

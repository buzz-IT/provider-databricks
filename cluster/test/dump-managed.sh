#!/usr/bin/env bash
# Best-effort dump of e2e managed resources and provider logs.
set -uo pipefail
KUBECTL="${KUBECTL:-kubectl}"

echo "=== managed resources ==="
${KUBECTL} get catalog,schema,secretscope,directory -A || true

echo "=== catalog describe ==="
${KUBECTL} describe catalog.unity.databricks.crossplane.io -A || true
${KUBECTL} describe catalog.unity.databricks.m.crossplane.io -A || true
${KUBECTL} describe schema.unity.databricks.crossplane.io -A || true
${KUBECTL} describe schema.unity.databricks.m.crossplane.io -A || true

echo "=== provider logs (last 200) ==="
${KUBECTL} -n "${CROSSPLANE_NAMESPACE:-crossplane-system}" logs -l pkg.crossplane.io/provider --tail=200 --all-containers=true || true
exit 0

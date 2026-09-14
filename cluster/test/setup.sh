#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
KUBECTL="${KUBECTL:-kubectl}"
CROSSPLANE_NAMESPACE="${CROSSPLANE_NAMESPACE:-crossplane-system}"
TEST_NAMESPACE="${UPTEST_NAMESPACE:-upbound-system}"

if [ -z "${UPTEST_CLOUD_CREDENTIALS:-}" ]; then
  UPTEST_CLOUD_CREDENTIALS="$("${ROOT}/cluster/test/render-credentials.sh")"
fi

echo "Creating namespaces..."
${KUBECTL} create namespace "${CROSSPLANE_NAMESPACE}" --dry-run=client -o yaml | ${KUBECTL} apply -f -
${KUBECTL} create namespace "${TEST_NAMESPACE}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Creating provider credential secrets..."
${KUBECTL} -n "${CROSSPLANE_NAMESPACE}" create secret generic provider-secret \
  --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" \
  --dry-run=client -o yaml | ${KUBECTL} apply -f -
${KUBECTL} -n "${TEST_NAMESPACE}" create secret generic provider-secret \
  --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" \
  --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 10m

echo "Waiting for provider pods..."
${KUBECTL} -n "${CROSSPLANE_NAMESPACE}" wait --for=condition=Available deployment --all --timeout=10m

echo "Creating cluster-scoped ProviderConfig..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: databricks.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: ${CROSSPLANE_NAMESPACE}
      key: credentials
EOF

echo "Creating namespaced ClusterProviderConfig..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: ${CROSSPLANE_NAMESPACE}
      key: credentials
EOF

echo "Creating namespaced ProviderConfig in ${TEST_NAMESPACE}..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: ${TEST_NAMESPACE}
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      key: credentials
EOF

#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"

PROJECT_MODULE="github.com/traefik/hub-crds"
KUBE_VERSION=v0.37.0
CONTROLLER_GEN_VERSION=v0.21.0
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd -P)"
BOILERPLATE="${SCRIPT_DIR}/boilerplate.go.tmpl"

cd "${REPO_ROOT}"

# Populate the module cache so kube_codegen.sh can be sourced from it.
# The other generators are installed by kube_codegen.sh itself.
go install "k8s.io/code-generator/cmd/deepcopy-gen@${KUBE_VERSION}"
go install "sigs.k8s.io/controller-tools/cmd/controller-gen@${CONTROLLER_GEN_VERSION}"

GOBIN_DIR="$(go env GOBIN)"
[ -n "${GOBIN_DIR}" ] || GOBIN_DIR="$(go env GOPATH)/bin"

CODEGEN_PKG="$(go env GOMODCACHE)/k8s.io/code-generator@${KUBE_VERSION}"
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "Generating deepcopy code ..."
kube::codegen::gen_helpers \
    --boilerplate "${BOILERPLATE}" \
    "${REPO_ROOT}/pkg/apis"

echo "Generating Hub clientset, listers and informers code ..."
kube::codegen::gen_client \
    --with-watch \
    --output-dir "${REPO_ROOT}/pkg/client" \
    --output-pkg "${PROJECT_MODULE}/pkg/client" \
    --boilerplate "${BOILERPLATE}" \
    "${REPO_ROOT}/pkg/apis"

echo "Generating the CRD definitions ..."
"${GOBIN_DIR}/controller-gen" \
    crd:crdVersions=v1,allowDangerousTypes=true \
    paths=./pkg/apis/hub/v1alpha1/... \
    output:dir=./pkg/apis/hub/v1alpha1/crd

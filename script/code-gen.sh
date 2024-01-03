#!/usr/bin/env bash
# shellcheck disable=SC2046

set -euo pipefail

source /go/src/k8s.io/code-generator/kube_codegen.sh

git config --global --add safe.directory /go/src/${PROJECT_MODULE}

kube::codegen::gen_helpers \
    --input-pkg-root "${PROJECT_MODULE}/pkg/apis" \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "/go/src/${PROJECT_MODULE}/script/boilerplate.go.tmpl"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root "${PROJECT_MODULE}/pkg/apis" \
    --output-pkg-root "${PROJECT_MODULE}/pkg/generated" \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "/go/src/${PROJECT_MODULE}/script/boilerplate.go.tmpl"

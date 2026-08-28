#!/usr/bin/env bash
# shellcheck disable=SC2046

set -euo pipefail

source /go/src/k8s.io/code-generator/kube_codegen.sh

git config --global --add safe.directory /go/src/${PROJECT_MODULE}

kube::codegen::gen_helpers \
    --boilerplate "/go/src/${PROJECT_MODULE}/script/boilerplate.go.tmpl" \
    "/go/src/${PROJECT_MODULE}/pkg/apis"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "/go/src/${PROJECT_MODULE}/pkg/client" \
    --output-pkg "${PROJECT_MODULE}/pkg/client" \
    --boilerplate "/go/src/${PROJECT_MODULE}/script/boilerplate.go.tmpl" \
    "/go/src/${PROJECT_MODULE}/pkg/apis"

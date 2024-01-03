#!/usr/bin/env bash
# shellcheck disable=SC2046

set -euo pipefail

PROJECT_MODULE="github.com/traefik/hub-crds"
IMAGE_NAME="kubernetes-codegen:latest"
CURRENT_DIR="$(pwd)"

echo "Building codegen Docker image..."
docker build --build-arg KUBE_VERSION=v0.28.4 \
             --build-arg USER="${USER}" \
             --build-arg UID="$(id -u)" \
             --build-arg GID="$(id -g)" \
             -f "./script/codegen.Dockerfile" \
             -t "${IMAGE_NAME}" \
             "."

echo "Generating Hub clientset, listers and informers code ..."
docker run --rm \
           -v "${CURRENT_DIR}:/go/src/${PROJECT_MODULE}" \
           -w "/go/src/${PROJECT_MODULE}" \
           -e "PROJECT_MODULE=${PROJECT_MODULE}" \
           "${IMAGE_NAME}" \
           bash ./script/code-gen.sh

echo "Generating the CRD definitions ..."
docker run --rm \
           -v "${CURRENT_DIR}:/go/src/${PROJECT_MODULE}" \
           -w "/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" \
           controller-gen crd:crdVersions=v1 \
           paths=./hub/v1alpha1/... \
           output:dir=./hub/v1alpha1/crd

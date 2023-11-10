# Traefik Hub CRDs

This repository contains the custom resource definitions (CRD) for [Traefik Hub](https://traefik.io/traefik-hub/).

- Group: `hub.traefik.io`
- Current version: `v1alpha1`

The YAML manifests can be found in `hub/v1alpha1/crd` and are built from `hub/v1alpha1/*.go`.

## Generate CRD manifests, client-sets, listers and informers

```shell
$> make build
```

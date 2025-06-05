# Traefik Hub CRDs

This repository contains the custom resource definitions (CRD) for [Traefik Hub](https://traefik.io/traefik-hub/).

Group: `hub.traefik.io`

Supported CRDs:

| Name                | Version  | 
|---------------------|----------|
| APIPortal           | v1alpha1 |
| APIPlan             | v1alpha1 |
| APICatalogItem      | v1alpha1 |
| APIBundle           | v1alpha1 |
| API                 | v1alpha1 |
| APIVersion          | v1alpha1 |
| ManagedSubscription | v1alpha1 |
| ManagedApplication  | v1alpha1 |

Deprecated CRDs:

| Name                | Version  | 
|---------------------|----------|
| AccessControlPolicy | v1alpha1 |
| APIRateLimit        | v1alpha1 |


The YAML manifests can be found in `hub/crd` and are built from `hub/v1alpha1/*.go`.

## Generate CRD manifests, client-sets, listers and informers

```shell
$> make build
```

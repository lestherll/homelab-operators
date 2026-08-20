# homelab-operators

Kubernetes controllers for my homelab cluster. One Go module, one binary, one
image — each operator is a controller inside a shared manager rather than a
deployment of its own.

Cluster manifests live in the [homelab](https://github.com/lestherll/homelab)
repo under `infrastructure/homelab-operators/`, applied by Flux.

## Controllers

| Controller | Group | Kind | Purpose |
| --- | --- | --- | --- |
| `janitor` | `janitor.homelab/v1alpha1` | `StalePodPolicy` | Deletes `Failed`/`Succeeded` pods past a TTL |

Nothing is implemented yet.

## Layout

```
api/<group>/<version>/          API types
internal/<group>/               Reconciliation logic, no Kubernetes client
internal/controller/<group>/    Reconcilers
cmd/main.go                     Manager setup and controller registry
config/                         Kustomize bases (kubebuilder-generated)
```

The reconcilers are deliberately thin. Anything worth testing lives in
`internal/<group>/` as pure functions, so the bulk of the test suite is table
tests rather than fake clients or a live API server.

## Running

```sh
make test                  # unit tests plus envtest
make run                   # against the cluster in ~/.kube/config
make manifests generate    # regenerate CRDs and deepcopy after type changes
```

Controllers can be selected with `--controllers=janitor,...`; the default is all
of them.

## Releasing

Tagging builds an image to `ghcr.io/lestherll/homelab-operators` and produces
`dist/install.yaml`. Vendor that into the homelab repo with the tag pinned.

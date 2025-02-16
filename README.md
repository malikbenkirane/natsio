# natsio demo

## Requirements

- `tofu` and `tofu init`
- a kubernetes cluster (we are using `kind` with [local registry](
https://kind.sigs.k8s.io/docs/user/local-registry/
))
- `nats` deployed to `nats` namespace (see also [nats helm repo](
https://docs.nats.io/running-a-nats-service/nats-kubernetes
))
  ```sh
  helm repo add nats https://nats-io.github.io/k8s/helm/charts/
  helm install --namespace nats nats nats/nats
  ```

## Build svctl locally

Build svctl container with the following steps
```sh
cd svctl
image=
podman build --tls-verify=false -t localhost:${regport}/natsio/svctl:[tag] .
podman push
```

Then update `image` variable before running `tofu apply`.

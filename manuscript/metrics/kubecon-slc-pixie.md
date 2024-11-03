# Metrics With Pixie

## Setup

> Install `px` by following the instructions at https://docs.px.dev/installing-pixie/install-schemes/cli

```bash
px deploy --cluster_name dot -y
```

## Do

* Open https://work.withpixie.ai in a browser

```sh
px scripts list

px scripts run px/namespaces

px scripts run px/namespace -- --help

px scripts run px/namespace -- --namespace production

px scripts show px/namespace

px live px/namespaces

# Press `Ctrl+c` to stop it

px live px/namespace -- --namespace production

# Press `Ctrl+c` to stop it
```

## Continue the Adventure

* [Jaeger](../tracing/kubecon-slc-jaeger.md)
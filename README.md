# kubestar

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/emcfarlane/kubestar?tab=doc)

Build kuberenetes config with [Starklark](https://github.com/google/starlark-go/blob/master/doc/spec.md) (a dialect of python) and [protobuffers](https://emcfarlane/starlarkproto).
Generate config with python functions, strongly typed to protobuffer objects.
Build out modules for different applications and import them with a function syntax for easy configuration.

Here's some example syntax to declare an apps/v1 Deployment object:
```python
apps_v1.Deployment(
  metadata={
    "name": "nginx-deployment",
    "labels": {
      "app": "nginx",
    },
  },
  spec=apps_v1.DeploymentSpec(
    replicas=3,
    selector=struct(matchLabels={
      "app": "nginx",
    }),
    template=core_v1.PodTemplateSpec(
      metadata=apis_meta_v1.ObjectMeta(
        labels={
          "app": "nginx",
        }),
      spec=core_v1.PodSpec(containers=[
        nginx_container(),
      ]),
    ),
  ),
)
```
Please see examples for more.

## How does it work?

1. `kubestar` loads all starlark files in a given glob pattern.
2. Each `main()` function is ran expecting an array of kubernetes protobuf objects.
4. Finally, the protobuf is than encoded to yaml and written to the named file with the '.star' -> '.yaml'.

Alternative to:
- [kustomize](https://github.com/kubernetes-sigs/kustomize)
- [skycfg](https://github.com/stripe/skycfg)

## Install

```
go install github.com/emcfarlane/kubestar@latest
```

## Examples

Have a look at the examples and try with:
```
kubestar examples/nginx_deployment.star -v 
```

```
examples
├── nginx_deployment.star
└── nginx_deployment.yaml (generated)
```

## Features

### Protobuf syntax

Protobuf wrapping is provided by [`github.com/emcfarlane/starlarkproto`](https://github.com/emcfarlane/starlarkproto). 
See the repo for more details. Similar types will be cast to protobuffer types.

### Load common modules

Load common modules with relative path syntax, or absolute based from the cmd init.
```python
load("./my_application.star", "object")
```

### Protobuffer file descriptors

Protobuffer file descriptors can be loaded using `proto.file(path)`.
See `protos.star` for all global namespaced file descriptors. 

```python
api_apps_v1 = proto.file("k8s.io/api/apps/v1/generated.proto")
```
Drop the `k8s.io/` prefix and `/generated.proto` suffix and finally replace `/` with `_`.

### Global config files

Important config can be declared in a global file with the flag `-global filename.star`. 
This will exec the file and expose as global config to all source files.

## Contrib

K8s protobuffers are currently gogo generated: https://github.com/kubernetes/kubernetes/issues/96564
We will need to clone [api](https://github.com/kubernetes/api) and [apimachinery](https://github.com/kubernetes/apimachinery) repos to a `k8s.io/` directory like:
```
k8s.io
├── api
└── apimachinery
```

Then generate the `protos.pb` file:
```sh
export GOOGLEAPIS_DIR=~/src/github.com/googleapis/api-common-protos
export K8S_DIR=~/src/github.com
./gen.sh
```

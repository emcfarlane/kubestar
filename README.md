# kubestar

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/emcfarlane/kubestar?tab=doc)

Build kuberenetes config with [Starklark](https://github.com/google/starlark-go/blob/master/doc/spec.md) (a subset dialect of python). 

```python
api_apps_v1.Deployment(
  metadata={
    "name": "nginx-deployment",
    "labels": {
      "app": "nginx",
    },
  },
  spec=api_apps_v1.DeploymentSpec(
    replicas=3,
    selector=struct(matchLabels={
      "app": "nginx",
    }),
    template=api_core_v1.PodTemplateSpec(
      metadata=apimachinery_pkg_apis_meta_v1.ObjectMeta(
        labels={
          "app": "nginx",
        }),
      spec=api_core_v1.PodSpec(containers=[
        nginx_container(),
      ]),
    ),
  ),
)
```
Please see examples.

## How does it work?

1. `kubestar` loads all starlark files in a given glob pattern.
2. Each `main()` function is ran expecting an array of kubernetes protobuf objects.
4. Finally, the protobuf is than encoded to yaml and written to the named file with the '.yaml' ext.

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

## Features

### Protobuf syntax

Protobuf wrapping is provided by [`github.com/emcfarlane/starlarkproto`](https://github.com/emcfarlane/starlarkproto). 

### Load common modules

Load common modules with relative path syntax, or absolute based from the cmd init.
```python
load("./relative.proto", "object")
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

Lets hack around this for now by scraping the protobuffers and regnerating it here.
```
protoc --go_out=source_relative:. $(find k8s.io/ -name "*.proto")
```

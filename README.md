# kubestar

Kubernetes configuration management, without all the YAML.
Build kubernetes objects with Starlark (a python-like dialect).
Based on [larking.io](https://larking.io) libraries.

![fun with yaml](./docs/yaml.jpeg)

## How does it work?

1. Kubestar loads all starlark files given.
2. Each `main()` function is ran expecting a kubernetes protobuf object returned.
3. Each file is optionally tested on any functions prefix with `test_...()`.
4. Finally, the protobuf is than encoded to yaml and written to the named file with the '.yaml' ext.

Alternative to:
- [kustomize](https://github.com/kubernetes-sigs/kustomize)
- [skycfg](https://github.com/stripe/skycfg)

## Install

```
go install github.com/emcfarlane/kubestar@latest
```

## Example

Have a look at the examples and try with:
```
kubestar examples/nginx_deployment.star
```

## Build

K8s protobuffers are currently gogo generated: https://github.com/kubernetes/kubernetes/issues/96564

Lets hack around this for now by scraping the protobuffers and regnerating it here.
```
protoc --go_out=source_relative:. $(find k8s.io/ -name "*.proto")
```

module github.com/emcfarlane/kubestar

go 1.18

replace (
	k8s.io/api => ./k8s.io/api
	k8s.io/apimachinery => ./k8s.io/apimachinery
)

require (
	github.com/emcfarlane/starlarkproto v0.0.0-20230424122739-88b5cbd6215b
	go.starlark.net v0.0.0-20230302034142-4b1e35fe2254
	google.golang.org/protobuf v1.30.0
	k8s.io/api v0.23.6
	k8s.io/apimachinery v0.23.6
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

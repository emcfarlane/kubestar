module github.com/emcfarlane/kubestar

go 1.18

replace (
	k8s.io/api => ./k8s.io/api
	k8s.io/apimachinery => ./k8s.io/apimachinery
)

require (
	github.com/emcfarlane/larking v0.0.0-20220412132153-bbedd19bbd7e
	github.com/emcfarlane/starlarkassert v0.0.0-20220406142958-771296b4bdb6
	go.starlark.net v0.0.0-20220328144851-d1966c6b9fcd
	google.golang.org/protobuf v1.28.0
	k8s.io/api v0.23.6
	k8s.io/apimachinery v0.23.6
	sigs.k8s.io/yaml v1.3.0
)

require (
	contrib.go.opencensus.io/integrations/ocsql v0.1.7 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/googleapis/gax-go/v2 v2.2.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	go.opencensus.io v0.23.0 // indirect
	gocloud.dev v0.25.0 // indirect
	golang.org/x/net v0.0.0-20220401154927-543a649e0bdd // indirect
	golang.org/x/sys v0.0.0-20220405210540-1e041c57c461 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.74.0 // indirect
	google.golang.org/genproto v0.0.0-20220401170504-314d38edb7de // indirect
	google.golang.org/grpc v1.45.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

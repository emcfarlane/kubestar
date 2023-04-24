package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emcfarlane/starlarkproto"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"sigs.k8s.io/yaml"
)

type k8Object struct {
	kind       string
	apiVersion string
}

var (
	//go:embed protos.star
	protoStar []byte

	//go:embed protos.pb
	protoPB []byte

	globals = starlark.StringDict{
		"struct": starlark.NewBuiltin("struct", starlarkstruct.Make),
		"proto":  starlarkproto.NewModule(protoregistry.GlobalFiles),
	}

	kindLooup = map[k8Object]protoreflect.FullName{}
)

func genAPIVersion(md protoreflect.FileDescriptor) string {
	return strings.ReplaceAll(
		strings.TrimPrefix(
			strings.TrimPrefix(
				string(md.FullName()),
				"k8s.io.api.",
			),
			"k8s.io.apimachinery.pkg.",
		),
		".", "/",
	)
}

func parseAPIVersion(apiVersion string) string {
	if strings.Count(apiVersion, "/") == 0 {
		return "core/" + apiVersion
	}
	return apiVersion
}

func init() {
	var descSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(protoPB, &descSet); err != nil {
		panic(err)
	}
	for _, file := range descSet.File {
		fd, err := protodesc.NewFile(file, protoregistry.GlobalFiles)
		if err != nil {
			panic(err)
		}

		if err := protoregistry.GlobalFiles.RegisterFile(fd); err != nil {
			panic(err)
		}

		apiVersion := genAPIVersion(fd)
		name := strings.ReplaceAll(apiVersion, "/", "_")

		globals[name] = starlarkproto.NewDescriptor(fd)

		//fmt.Printf("%s = proto.file(%q)\n", name, fd.Path())

		for i := 0; i < fd.Messages().Len(); i++ {
			md := fd.Messages().Get(i)
			k8 := k8Object{
				kind:       string(md.Name()),
				apiVersion: apiVersion,
			}
			kindLooup[k8] = md.FullName()
		}
	}
}

func yamlToProtos(src string) ([]proto.Message, error) {
	var msgs []proto.Message
	for _, part := range strings.Split(src, "---\n") {
		jb, err := yaml.YAMLToJSON([]byte(part))
		if err != nil {
			return nil, err
		}
		fmt.Println(string(jb))

		var m map[string]interface{}
		if err := json.Unmarshal(jb, &m); err != nil {
			return nil, err
		}
		kind, ok := m["kind"].(string)
		if !ok {
			return nil, fmt.Errorf("kind invalid")
		}
		apiVersion, ok := m["apiVersion"].(string)
		if !ok {
			return nil, fmt.Errorf("apiVersion invalid")
		}
		apiVersion = parseAPIVersion(apiVersion)

		delete(m, "kind")
		delete(m, "apiVersion")
		jb, err = json.Marshal(m)
		if err != nil {
			return nil, err
		}

		typ, ok := kindLooup[k8Object{
			kind:       kind,
			apiVersion: apiVersion,
		}]
		if !ok {
			return nil, fmt.Errorf("kind %s not found", kind)
		}
		desc, err := protoregistry.GlobalFiles.FindDescriptorByName(typ)
		if err != nil {
			return nil, fmt.Errorf("type %s not found: %w", typ, err)
		}
		md := desc.(protoreflect.MessageDescriptor)
		msg := dynamicpb.NewMessage(md)
		if err := protojson.Unmarshal(jb, msg); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)

		//fmt.Printf("kind: %s, apiVersion: %s\n", kind, apiVersion)
	}
	return msgs, nil
}

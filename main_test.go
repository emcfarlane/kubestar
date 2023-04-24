package main_test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var flagGenStar = flag.Bool("genstar", false, "Verbose mode.")

func TestProto(t *testing.T) {
	flag.Parse()

	var b strings.Builder
	b.WriteString("# generated by go generate; DO NOT EDIT\n")
	protoregistry.GlobalFiles.RangeFiles(func(p protoreflect.FileDescriptor) bool {
		if name := string(p.FullName()); strings.HasPrefix(name, "k8s.io") {
			name = strings.TrimPrefix(name, "k8s.io.")
			name = strings.ReplaceAll(name, ".", "_")

			fmt.Fprintf(&b, "%s = proto.file(%q)\n", name, p.Path())
		}
		return true
	})
	if *flagGenStar {
		if err := os.WriteFile("protos.star", []byte(b.String()), 0644); err != nil {
			t.Fatal(err)
		}
		t.Log("wrote protos.star")
	} else {
		t.Log(b.String())
	}
}

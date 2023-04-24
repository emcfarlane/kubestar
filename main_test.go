package main_test

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"testing"
)

func TestProto(t *testing.T) {
	fileSystem := os.DirFS("./k8s.io")

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() || !strings.HasSuffix(path, ".proto") {
			return nil
		}
		name := strings.TrimSuffix(path, "/generated.proto")
		name = strings.ReplaceAll(name, "/", "_")

		fmt.Printf("%s = proto.file(%q)\n", name, "k8s.io/"+path)
		return nil
	})
}

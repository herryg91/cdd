package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	descriptor "github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	crudgenerator "github.com/herryg91/cdd/protoc-gen-cdd/generator/crud-generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	fGoModuleName = flag.String("go-module-name", "", "Go module name, check in go.mod file. This needed for local import prefix. example: github.com/herryg91/cdd/examples/province-api")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		registry := descriptor.New(*plugin.Request)
		if *fGoModuleName == "" {
			return fmt.Errorf("Option `go-module-name` is required. Example `--cdd_opt go-module-name=$(go list -m)>`")
		}
		gen := crudgenerator.New(registry, *fGoModuleName)

		if gen != nil {
			files, err := gen.Generate()
			if err != nil {
				return err
			}
			for _, f := range files {

				genFile := plugin.NewGeneratedFile(f.Filename, f.GoImportPath)
				if _, err := genFile.Write([]byte(f.Content)); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

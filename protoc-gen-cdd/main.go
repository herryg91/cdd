package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	descriptor "github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	crudgenerator "github.com/herryg91/cdd/protoc-gen-cdd/generator/crud-generator"
	grstframework "github.com/herryg91/cdd/protoc-gen-cdd/generator/grst-framework"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	fType = flag.String("type", "grst", "option: grst|crud")
	/*grst specific options*/
	fProtocGoOut = flag.Bool("protoc-gen-go", true, "generate *.pb.go (calling `protoc-gen-go`) with additional features, such as request validation & default value. protoc-gen-go version: v1.25.0. default: true")
	/*crud specific options*/
	fImportPath = flag.String("go-import-path", "", "go import path. example: github.com/herryg91/cdd/examples/province-api")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		registry := descriptor.New(*plugin.Request)
		var gen generator.Generator = nil
		switch *fType {
		case "grst":
			gen = grstframework.New(registry, *plugin, *fProtocGoOut)
		case "crud":
			if *fImportPath == "" {
				return fmt.Errorf("Option `go-import-path` is required. Example `--cdd_opt go-import-path=$(go list -m)>`")
			}
			gen = crudgenerator.New(registry, *fImportPath)
		default:
			return fmt.Errorf("Invalid option `type`, got: %s, expect: %s", *fType, "grst|crud")
		}
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

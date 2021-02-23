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
	fType        = flag.String("type", "grst", "option: grst|grst-validation|crud")
	fProtocGoOut = flag.Bool("protoc-gen-go", true, "generate *.pb.go (calling `protoc-gen-go`) with additional features, such as request validation & default value. protoc-gen-go version: v1.25.0. default: true")
	fImportPath  = flag.String("go-import-path", "", "go import path. example: github.com/herryg91/cdd/examples/province-api")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		// log.Println("Parsing code generator request")

		registry := descriptor.New(*plugin.Request)

		var gen generator.Generator = nil
		switch *fType {
		case "grst":
			gen = grstframework.New(registry, *plugin, *fProtocGoOut)
		case "crud":
			if *fImportPath == "" {
				return fmt.Errorf("Parameter `pkgpath` for `crud` is required, found: %s", *fImportPath)
			}
			gen = crudgenerator.New(registry, *fImportPath)
		default:
			return fmt.Errorf("invalid opt `type`, got: %s, expect: %s", *fType, "grst|grst-validation|crud")
		}
		if gen != nil {
			files, err := gen.Generate()
			if err != nil {
				return err
			}
			// log.Println("Generating...")
			for _, f := range files {

				genFile := plugin.NewGeneratedFile(f.Filename, f.GoImportPath)
				if _, err := genFile.Write([]byte(f.Content)); err != nil {
					return err
				}
				// log.Printf("NewGeneratedFile: %s\n", f.Filename)
			}
		}

		return nil
	})
}

package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	descriptor "github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	grstframework "github.com/herryg91/cdd/protoc-gen-cdd/generator/grst-framework"
	scaffold_mysql "github.com/herryg91/cdd/protoc-gen-cdd/generator/scaffold-mysql"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	fType = flag.String("type", "grst", "option: grst|scaffold-mysql")
	/*grst specific options*/
	fProtocGoOut = flag.Bool("protoc-gen-go", true, "generate *.pb.go (calling `protoc-gen-go`) with additional features, such as request validation & default value. protoc-gen-go version: v1.25.0. default: true")
	/*scaffold-mysql specific options*/
	fGoModuleName = flag.String("go-module-name", "", "Go module name, check in go.mod file. This needed for local import prefix. example: github.com/herryg91/cdd/examples/province-api")
	fVersion      = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if *fVersion {
		fmt.Println("protoc-gen-cdd v1.0.0")
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		registry := descriptor.New(*plugin.Request)
		var gen generator.Generator = nil
		switch *fType {
		case "grst":
			gen = grstframework.New(registry, *plugin, *fProtocGoOut)
		case "scaffold-mysql":
			if *fGoModuleName == "" {
				return fmt.Errorf("Option `go-module-name` is required. Example `--cdd_opt go-module-name=$(go list -m)>`")
			}
			gen = scaffold_mysql.New(registry, *fGoModuleName)
		default:
			return fmt.Errorf("Invalid option `type`, got: %s, expect: %s", *fType, "grst|scaffold-mysql")
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

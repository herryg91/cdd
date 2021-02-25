package main

import (
	"flag"

	"github.com/golang/glog"
	descriptor "github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	grstframework "github.com/herryg91/cdd/protoc-gen-cdd/generator/grst-framework"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	fProtocGoOut = flag.Bool("protoc-gen-go", true, "generate *.pb.go (calling `protoc-gen-go`) with additional features, such as request validation & default value. protoc-gen-go version: v1.25.0. default: true")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		registry := descriptor.New(*plugin.Request)
		gen := grstframework.New(registry, *plugin, *fProtocGoOut)
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

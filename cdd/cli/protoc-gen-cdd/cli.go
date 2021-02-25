package protocgencdd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/herryg91/cdd/cdd/cli/protoc"
)

type ProtocGenCdd struct {
}

func NewProtocGenCdd() *ProtocGenCdd {
	return &ProtocGenCdd{}
}

func (pgc *ProtocGenCdd) GenerateGrst(protoFilename string, inputPath string, outputPath string) error {
	outputPath = outputPath + strings.Replace(filepath.Base(protoFilename), filepath.Ext(protoFilename), "", -1)
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtoPath("grpc/proto/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "grst"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "go-grpc", Opts: map[string]string{}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "grpc-gateway", Opts: map[string]string{"logtostderr": "true", "generate_unbound_methods": "true"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	fmt.Println("Run Protoc 1 =>", "in="+inputPath+" | out="+outputPath)

	err := p.Exec(filepath.Base(protoFilename), true)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateCrud(protoFilename string, inputPath string, outputPath string, goModuleName string) error {
	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "crud", "go-module-name": goModuleName}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	fmt.Println("Run Protoc 3 =>")
	err := p.Exec(filepath.Base(protoFilename), true)
	if err != nil {
		return err
	}
	return nil
}

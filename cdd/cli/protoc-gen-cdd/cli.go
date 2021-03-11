package protocgencdd

import (
	"log"
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

func (pgc *ProtocGenCdd) GenerateGrst(protoFilename string, inputPath string, outputPath string, printLog bool) error {
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

	// fmt.Println("Run protoc-gen-cdd type=grst", "| in = "+inputPath+" | out ="+outputPath)
	log.Println("Generating file [type=grst]: " + inputPath + " | out" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateScaffoldMysql(protoFilename string, inputPath string, outputPath string, goModuleName string, printLog bool) error {
	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "scaffold-mysql", "go-module-name": goModuleName}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	// fmt.Println("Run protoc-gen-cdd type=scaffold-mysql", "| in = "+inputPath+" | out ="+outputPath)
	log.Println("Generating file [type=scaffold-mysql]: " + inputPath + " | out" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

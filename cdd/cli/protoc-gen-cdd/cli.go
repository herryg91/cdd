package protocgencdd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/herryg91/cdd/cdd/cli/protoc"
	"github.com/herryg91/cdd/cdd/pkg/configFile"
	"github.com/herryg91/cdd/cdd/pkg/helpers"
)

type ProtocGenCdd struct {
	ExtProtoPaths []string
}

func NewProtocGenCdd() *ProtocGenCdd {
	extProtoPaths := []string{
		"$HOME/.cdd/ext/cddapis/",
		"$HOME/.cdd/ext/googleapis/",
	}

	homeDir := helpers.HomeDir()

	if _, err := os.Stat(homeDir + "/.cdd/config"); err == nil {
		body, err := os.ReadFile(homeDir + "/.cdd/config")
		if err == nil {
			bodyObj := configFile.ConfigFile{}
			json.Unmarshal(body, &bodyObj)

			extProtoPaths = bodyObj.Extensions
		}
	}

	return &ProtocGenCdd{
		ExtProtoPaths: extProtoPaths,
	}
}

func (pgc *ProtocGenCdd) GenerateGrst(protoFilename string, inputPath string, outputPath string, printLog bool) error {
	outputPath = outputPath + strings.Replace(filepath.Base(protoFilename), filepath.Ext(protoFilename), "", -1)
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	// p.AddProtoPath("$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis")
	for _, pp := range pgc.ExtProtoPaths {
		p.AddProtoPath(pp)
	}
	// p.AddProtoPath("$HOME/.cdd/ext/cddapis/")
	// p.AddProtoPath("$HOME/.cdd/ext/googleapis/")
	// p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	// p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "grst"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "go-grpc", Opts: map[string]string{}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "grpc-gateway", Opts: map[string]string{"logtostderr": "true", "generate_unbound_methods": "true"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=grst]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateMysqlModel(protoFilename string, inputPath string, outputPath string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "mysql-model"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=mysql-model]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateEntity(protoFilename string, inputPath string, outputPath string, entities []string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "entity", "name": strings.Join(entities, "|")}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=entity]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateUsecaseMysql(protoFilename string, inputPath string, outputPath string, modelName string, goModuleName string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "usecase-mysql", "name": modelName, "go-module-name": goModuleName}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=usecase-mysql]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

package gen

import (
	"path/filepath"

	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	gocli "github.com/herryg91/cdd/cdd/cli/go"
	protocgencdd "github.com/herryg91/cdd/cdd/cli/protoc-gen-cdd"
	"github.com/herryg91/cdd/cdd/cli/serviceYaml"
	"github.com/spf13/cobra"
)

type GenGoCmd struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	serviceYamlFile string
}

func NewGenGoCmd() *GenGoCmd {
	c := &GenGoCmd{
		Command: &cobra.Command{
			Use:   "go",
			Short: "generate cdd framework",
			Long:  "generate cdd framework",
			Run:   nil,
		},
		protocGenCddCli: protocgencdd.NewProtocGenCdd(),
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVar(&c.serviceYamlFile, "service-yaml", "service.yaml", "service.yaml file path")
	return c
}

type ContractToGenerate struct {
	protoInput             string
	outputGrstDir          string
	outputScaffoldMysqlDir string
	scaffoldMysql          bool
}

func (c *GenGoCmd) runCommand(cmd *cobra.Command, args []string) error {
	svcYaml, err := serviceYaml.GetServiceYAML(c.serviceYamlFile)
	if err != nil {
		return err
	}
	outputGrst := svcYaml.Contract.Config.OutputGrst
	outputScaffoldMysql := svcYaml.Contract.Config.OutputScaffoldMysql
	outputDependency := svcYaml.Contract.Config.OutputDependency
	if outputGrst == "" {
		outputGrst = "grpc/pb/"
	}
	if outputDependency == "" {
		outputDependency = "grpc/pb-deps/"
	}

	contractsToGenerate := []ContractToGenerate{}
	// Setup proto contract for main service
	for _, file := range svcYaml.Contract.ProtoFiles {
		contractsToGenerate = append(contractsToGenerate, ContractToGenerate{
			protoInput:             file,
			outputGrstDir:          outputGrst,
			outputScaffoldMysqlDir: outputScaffoldMysql,
			scaffoldMysql:          true})
	}

	// Setup proto contract for dependencies services
	for _, svcFilePath := range svcYaml.Dependencies.Services {
		svcYamlDependency, err := serviceYaml.GetServiceYAML(svcFilePath)
		if err != nil {
			return err
		}

		dirDependency, _ := filepath.Split(svcFilePath)

		for _, file := range svcYamlDependency.Contract.ProtoFiles {
			contractsToGenerate = append(contractsToGenerate, ContractToGenerate{
				protoInput:             dirDependency + "/" + file,
				outputGrstDir:          outputDependency,
				outputScaffoldMysqlDir: outputScaffoldMysql,
				scaffoldMysql:          false})
		}
	}
	// generate grpc pb
	uiprogress.Start()
	bar := uiprogress.AddBar(len(contractsToGenerate)).AppendCompleted().PrependElapsed()
	bar.Width = 50
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		if b.Current() < len(contractsToGenerate) && b.Current() >= 0 {
			return strutil.Resize("Generating "+contractsToGenerate[b.Current()].protoInput, 30)
		} else {
			return strutil.Resize("done", 30)
		}
	})
	for _, ctg := range contractsToGenerate {
		currentModule, err := gocli.GetCurrentModule()
		if err != nil {
			return err
		}

		dir, filename := filepath.Split(ctg.protoInput)
		err = c.protocGenCddCli.GenerateGrst(filename, dir, ctg.outputGrstDir)
		if err != nil {
			return err
		}

		if ctg.scaffoldMysql {
			err = c.protocGenCddCli.GenerateScaffoldMysql(filename, dir, ctg.outputScaffoldMysqlDir, currentModule)
			if err != nil {
				return err
			}
		}
		bar.Incr()
	}
	bar.Incr()

	return nil
}

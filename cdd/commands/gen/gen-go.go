package gen

import (
	"path/filepath"

	protocgencdd "github.com/herryg91/cdd/cdd/cli/protoc-gen-cdd"
	"github.com/herryg91/cdd/cdd/pkg/serviceYaml"
	"github.com/spf13/cobra"
)

type GenGoCmd struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	serviceYamlFile string
	printLog        bool
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
	c.Command.Flags().BoolVar(&c.printLog, "print", false, "print log")
	return c
}

type ContractToGenerate struct {
	protoInput          string
	outputGrstDir       string
	outputMysqlModelDir string
	mysqlModel          bool
}

const defaultOutputMysqlModel = "drivers/datasource/mysql/"
const defaultOutputGrst = "drivers/handler/grst/"
const defaultOutputDependency = "drivers/external/grst/"

func (c *GenGoCmd) runCommand(cmd *cobra.Command, args []string) error {
	svcYaml, err := serviceYaml.GetServiceYAML(c.serviceYamlFile)
	if err != nil {
		return err
	}
	outputGrst := svcYaml.Contract.Config.OutputGrst
	outputMysqlModel := svcYaml.Contract.Config.OutputMysqlModel
	outputDependency := svcYaml.Contract.Config.OutputDependency
	if outputGrst == "" {
		outputGrst = defaultOutputGrst
	}
	if outputDependency == "" {
		outputDependency = defaultOutputDependency
	}
	if outputMysqlModel == "" {
		outputMysqlModel = defaultOutputMysqlModel
	}

	contractsToGenerate := []ContractToGenerate{}
	// Setup proto contract for main service
	for _, file := range svcYaml.Contract.ProtoFiles {
		contractsToGenerate = append(contractsToGenerate, ContractToGenerate{
			protoInput:          file,
			outputGrstDir:       outputGrst,
			outputMysqlModelDir: outputMysqlModel,
			mysqlModel:          true,
		})
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
				protoInput:          dirDependency + "/" + file,
				outputGrstDir:       outputDependency,
				outputMysqlModelDir: outputMysqlModel,
				mysqlModel:          false,
			})
		}
	}
	// generate grpc pb
	for _, ctg := range contractsToGenerate {
		dir, filename := filepath.Split(ctg.protoInput)
		err = c.protocGenCddCli.GenerateGrst(filename, dir, ctg.outputGrstDir, c.printLog)
		if err != nil {
			return err
		}

		// if ctg.mysqlModel {
		// 	err = c.protocGenCddCli.GenerateMysqlModel(filename, dir, ctg.outputMysqlModelDir, c.printLog)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
}

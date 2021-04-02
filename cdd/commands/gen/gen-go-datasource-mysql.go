package gen

import (
	"fmt"
	"path/filepath"

	ds_mysql_gen "github.com/herryg91/cdd/cdd/pkg/gen-go-datasource-mysql"
	"github.com/spf13/cobra"
)

type GenGoDsMysql struct {
	Command   *cobra.Command
	inputFile string
	modelName string
	outputDir string
}

func NewGenGoDatasourceMysqlCmd() *GenGoDsMysql {
	c := &GenGoDsMysql{
		Command: &cobra.Command{
			Use:     "go-ds-mysql",
			Aliases: []string{"go-dsm"},
			Short:   "generate query mysql on datasources based on struct model",
			Long:    "generate query mysql on datasources based on struct model",
			Run:     nil,
		},
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.inputFile, "input", "f", "", "input file path (*.go file)")
	c.Command.Flags().StringVarP(&c.modelName, "model-name", "n", "", "struct model name (case insensitive)")
	c.Command.Flags().StringVarP(&c.outputDir, "output-dir", "o", "", "output directory (optional). Default: same with input file's directory")
	return c
}

func (c *GenGoDsMysql) runCommand(cmd *cobra.Command, args []string) error {
	if c.inputFile == "" {
		return fmt.Errorf("--input or -f is required")
	}
	if c.outputDir == "" {
		c.outputDir = filepath.Dir(c.inputFile)
	}

	err := ds_mysql_gen.Generate(c.inputFile, c.modelName, c.outputDir)
	return err
}

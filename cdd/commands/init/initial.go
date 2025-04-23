package init

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/herryg91/cdd/cdd/pkg/configFile"
	"github.com/herryg91/cdd/cdd/pkg/helpers"
	"github.com/spf13/cobra"
)

type InitCmd struct {
	*cobra.Command
}

func NewInitCmd() *InitCmd {
	c := &InitCmd{}
	c.Command = &cobra.Command{
		Use:   "init",
		Short: "Initiate configuration",
		Long:  "Initiate configuration",
	}
	c.Command.RunE = c.runCommand
	return c
}

func (c *InitCmd) runCommand(cmd *cobra.Command, args []string) error {
	cf := configFile.ConfigFile{
		Extensions: []string{
			"$HOME/.cdd/ext/cddapis/",
			"$HOME/.cdd/ext/googleapis/",
		},
	}
	current_data, _ := json.MarshalIndent(cf, "", "    ")
	homeDir := helpers.HomeDir()

	if _, err := os.Stat(homeDir + "/.cdd"); !os.IsNotExist(err) {
		os.Mkdir(homeDir+"/.cdd", os.ModePerm)
	}

	if _, err := os.Stat(homeDir + "/.cdd/config"); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("Config file is exist")
		return nil
	}

	return os.WriteFile(homeDir+"/.cdd/config", current_data, 0644)
}

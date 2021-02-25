package install

import (
	"github.com/spf13/cobra"
)

type InstallCmd struct {
	*cobra.Command
}

func NewInstallCmd() *InstallCmd {
	c := &InstallCmd{}
	c.Command = &cobra.Command{
		Use:   "install",
		Short: "install generator requirement",
		Long:  "Installing CDD generator requirements",
	}
	return c
}

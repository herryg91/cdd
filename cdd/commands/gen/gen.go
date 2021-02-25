package gen

import (
	"github.com/spf13/cobra"
)

type GenCmd struct {
	*cobra.Command
}

func NewGenCmd() *GenCmd {
	c := &GenCmd{}
	c.Command = &cobra.Command{
		Use:   "gen",
		Short: "faraday framework x grst codegen tools",
		Long:  "command to generate code faraday framework (grst)",
	}
	c.AddCommand(NewGenGoCmd().Command)
	return c
}

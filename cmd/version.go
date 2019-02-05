package cmd

import (
	"fmt"

	"github.com/cbush06/kosher/common"
	"github.com/spf13/cobra"
)

type versionCommand struct {
	name    string
	command *cobra.Command
}

var cmdVersion = &versionCommand{
	name: "version",
	command: &cobra.Command{
		Use:   "version",
		Short: "displays the version of Kosher running",
		Long:  `version displays the version of Kosher running`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(common.CurrentVersion.Version())
			return nil
		},
	},
}

func (v *versionCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(v.command)
}

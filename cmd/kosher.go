package cmd

import (
	"github.com/spf13/cobra"
)

type kosherCommand struct {
	name    string
	command *cobra.Command
}

func buildKosherCommand() *kosherCommand {
	return &kosherCommand{
		name: "kosher",
		command: &cobra.Command{
			Use:   "kosher",
			Short: "main command for running Kosher",
			Long: `kosher is the main command, used to run Kosher.
			
	Kosher is a simple and powerful tool for powering behavior-driven
	development methodology by automating functional testing.
			
	Complete documentation is available at https://cbush06.github.io/kosher/.`,
			RunE: func(cmd *cobra.Command, args []string) error {
				cmd.Help()
				return nil
			},
		},
	}
}

func (k *kosherCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(k.command)
}

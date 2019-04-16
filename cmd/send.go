package cmd

import (
	"github.com/spf13/cobra"
)

type sendCommand struct {
	name    string
	command *cobra.Command
}

var cmdSend = &sendCommand{
	name: "send",
	command: &cobra.Command{
		Use:   "send",
		Short: "sends results to a remote system",
		Long:  `send transmits the results stored in a Kosher-generated Cucumber JSON file to a remote system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	},
}

func (s *sendCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(s.command)
}

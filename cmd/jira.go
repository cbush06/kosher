package cmd

import (
	"log"
	"os"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/integrations"
	"github.com/spf13/cobra"
)

type jiraCommand struct {
	name    string
	command *cobra.Command
}

var cmdJira = &jiraCommand{
	name: "jira",
	command: &cobra.Command{
		Use:   "jira",
		Short: "sends results to a Jira system creating tickets for each failed scenario",
		Long:  `jira creates a new Jira ticket for each failed scenario. The fields of the ticket (e.g. type, labels, summary, description, etc.) may be customized via the settings.json file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// determine where the executable was called from
			workingDir, _ := os.Getwd()
			if fileSys, err = fs.NewFs(workingDir); err != nil {
				log.Fatal(err)
			}

			// build the settings file based on the working directory
			settings = config.NewSettings(fileSys)

			if err := integrations.SendTo(integrations.Jira, settings); err != nil {
				log.Fatalln(err)
			}

			return nil
		},
	},
}

func (s *jiraCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(s.command)
}

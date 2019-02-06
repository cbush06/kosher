package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"

	"github.com/spf13/cobra"
)

type runCommand struct {
	name    string
	command *cobra.Command
}

var cmdRun = &runCommand{
	name: "run",
	command: &cobra.Command{
		Use:   "run",
		Short: "executes your tests",
		Long:  `run executes your tests. Depending on the arguments provided, it may execute all tests, a specific test, or tests in one or more subdirectories.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, arg []string) error {
			var path string
			if len(arg) < 1 {
				path, _ = os.Getwd()
			} else {
				path, _ = filepath.Abs(filepath.Clean(arg[0]))
			}

			fs, err := fs.NewFs(path)

			if err != nil {
				return err
			}

			settings := config.NewSettings(fs)

			if settings.Environments == nil {
				log.Fatal("Error")
			}

			return nil
		},
	},
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(r.command)
}

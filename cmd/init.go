package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kosher/kosherfs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type initCommand struct {
	name    string
	command *cobra.Command
}

var initEmpty bool
var initForce bool

var cmdInit = &initCommand{
	name: "init",
	command: &cobra.Command{
		Use:   "init [path to directory]",
		Short: "Initializes a new Kosher project",
		Long:  `init creates the necessary project structure with simple example tests and config files to quickly get you started.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			path := ""
			if len(args) < 1 {
				if path, err = os.Getwd(); err != nil {
					log.Fatal(err)
				}
			} else {
				if path, err = filepath.Abs(filepath.Clean(args[0])); err != nil {
					log.Fatal(err)
				}
			}

			return initProject(kosherfs.NewDefault(viper.New()), path, initForce)
		},
	},
}

func (i *initCommand) registerWith(cmd *cobra.Command) {
	i.command.Flags().BoolVarP(&initForce, "force", "f", false, "Create a project inside a non-empty directory.")
	i.command.Flags().BoolVarP(&initEmpty, "empty", "e", false, "Create an empty project (no example tests).")
	cmd.AddCommand(i.command)
}

func initProject(fs *kosherfs.Fs, basepath string, force bool) error {
	projectStructure := []string{
		filepath.Join(basepath, "features"),
		filepath.Join(basepath, "config"),
	}
}

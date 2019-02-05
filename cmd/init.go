package cmd

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/resources"
	"github.com/spf13/cobra"
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

			return initProject(fs.NewFs(path), path, initForce)
		},
	},
}

func (i *initCommand) registerWith(cmd *cobra.Command) {
	i.command.Flags().BoolVarP(&initForce, "force", "f", false, "Create a project inside a non-empty directory.")
	i.command.Flags().BoolVarP(&initEmpty, "empty", "e", false, "Create an empty project (no example tests).")
	cmd.AddCommand(i.command)
}

// initProject initializes a new Kosher project with configuration file templates and a sample feature file
func initProject(fs *fs.Fs, basepath string, force bool) error {
	var (
		featuresDir      = filepath.Join(basepath, common.FeaturesDir)
		configDir        = filepath.Join(basepath, common.ConfigDir)
		environmentsJSON = filepath.Join(configDir, common.EnvironmentsFile)
		pagesJSON        = filepath.Join(configDir, common.PagesFile)
		selectorsJSON    = filepath.Join(configDir, common.SelectorsFile)
		settingsJSON     = filepath.Join(configDir, common.SettingsFile)
		exampleFeature   = filepath.Join(featuresDir, "example.feature")
	)

	projectStructure := []string{
		featuresDir,
		configDir,
	}

	projectFiles := []string{
		environmentsJSON,
		pagesJSON,
		selectorsJSON,
		settingsJSON,
		exampleFeature,
	}

	if exists, _ := afero.Exists(fs.WorkingDir, basepath); exists {
		if isDir, _ := afero.IsDir(fs.WorkingDir, basepath); !isDir {
			return errors.New(basepath + " exists but is not a directory...aborting...")
		}

		isEmpty, _ := afero.IsEmpty(fs.WorkingDir, basepath)

		switch {
		case !isEmpty && !force:
			return errors.New(basepath + " already exists and is not empty...aborting...")

		case !isEmpty && force:
			all := append(projectStructure, projectFiles...)
			for _, path := range all {
				if pathExists, _ := afero.Exists(fs.WorkingDir, path); pathExists {
					return errors.New(path + " already exists...aborting...")
				}
			}
		}
	}

	for _, dir := range projectStructure {
		if err := fs.WorkingDir.MkdirAll(dir, 0777); err != nil {
			return errors.New("Failed to create dir " + dir)
		}
	}

	afero.WriteReader(fs.WorkingDir, settingsJSON, bytes.NewBufferString(resources.GetSettingsJSON()))
	afero.WriteReader(fs.WorkingDir, environmentsJSON, bytes.NewBufferString(resources.GetEnvironmentsJSON()))
	afero.WriteReader(fs.WorkingDir, pagesJSON, bytes.NewBufferString(resources.GetPagesJSON()))
	afero.WriteReader(fs.WorkingDir, selectorsJSON, bytes.NewBufferString(resources.GetSelectorsJSON()))

	return nil
}

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/afero"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/resources/configfiles"
	"github.com/spf13/cobra"
)

type initCommand struct {
	name    string
	command *cobra.Command
}

var (
	validPlatforms = []string{"desktop", "web"}
	initPlatform   = "web"
	initEmpty      bool
	initForce      bool
)

var cmdInit = &initCommand{
	name: "init",
	command: &cobra.Command{
		Use:   "init [flags] [project name] [path to directory]",
		Short: "Initializes a new Kosher project",
		Long:  `init creates the necessary project structure with simple example tests and config files to quickly get you started.`,
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			path := ""
			if len(args) < 2 {
				if path, err = os.Getwd(); err != nil {
					log.Fatal(err)
				}
			} else {
				if path, err = filepath.Abs(filepath.Clean(args[1])); err != nil {
					log.Fatal(err)
				}
			}

			if platformIndex := sort.SearchStrings(validPlatforms, initPlatform); platformIndex >= len(validPlatforms) || validPlatforms[platformIndex] != initPlatform {
				log.Fatal(errors.New("Invalid platform specified [" + initPlatform + "]. Valid options are: " + strings.Join(validPlatforms, ", ")))
			}

			if err = initProject(args[0], &afero.OsFs{}, path, initForce, initEmpty); err == nil {
				fmt.Printf("Project [%s] initialized...\n", args[0])
			}

			return err
		},
	},
}

func (i *initCommand) registerWith(cmd *cobra.Command) {
	i.command.Flags().StringVarP(&initPlatform, "platform", "p", "web", "Set the platform: "+strings.Join(validPlatforms, ", "))
	i.command.Flags().BoolVarP(&initForce, "force", "f", false, "Create a project inside a non-empty directory.")
	i.command.Flags().BoolVarP(&initEmpty, "empty", "e", false, "Create an empty project (no example tests).")
	cmd.AddCommand(i.command)
}

// initProject initializes a new Kosher project with configuration file templates and a sample feature file
func initProject(projectName string, fs *afero.OsFs, basepath string, force bool, empty bool) error {
	var (
		featuresDir      = filepath.Join(basepath, common.FeaturesDir)
		configDir        = filepath.Join(basepath, common.ConfigDir)
		resultsDir       = filepath.Join(basepath, common.ResultsDir)
		environmentsJSON = filepath.Join(configDir, common.EnvironmentsFile)
		pagesJSON        = filepath.Join(configDir, common.PagesFile)
		selectorsJSON    = filepath.Join(configDir, common.SelectorsFile)
		settingsJSON     = filepath.Join(configDir, common.SettingsFile)
		exampleFeature   = filepath.Join(featuresDir, common.ExampleFeatureFile)
	)

	projectStructure := []string{
		featuresDir,
		configDir,
		resultsDir,
	}

	projectFiles := []string{
		environmentsJSON,
		pagesJSON,
		selectorsJSON,
		settingsJSON,
		exampleFeature,
	}

	if exists, _ := afero.Exists(fs, basepath); exists {
		if isDir, _ := afero.IsDir(fs, basepath); !isDir {
			return errors.New(basepath + " exists but is not a directory...aborting...")
		}

		isEmpty, _ := afero.IsEmpty(fs, basepath)

		switch {
		case !isEmpty && !force:
			return errors.New(basepath + " already exists and is not empty...aborting...")

		case !isEmpty && force:
			all := append(projectStructure, projectFiles...)
			for _, path := range all {
				if pathExists, _ := afero.Exists(fs, path); pathExists {
					return errors.New(path + " already exists...aborting...")
				}
			}
		}
	}

	for _, dir := range projectStructure {
		if err := fs.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("Failed to create dir [%s]: %s", dir, err)
		}
	}

	if initPlatform == "web" {
		afero.WriteReader(fs, settingsJSON, bytes.NewBufferString(configfiles.GetSettingsJSON(projectName)))
	} else if initPlatform == "desktop" {
		afero.WriteReader(fs, settingsJSON, bytes.NewBufferString(configfiles.GetSettingsDesktopJSON(projectName)))
	}

	afero.WriteReader(fs, environmentsJSON, bytes.NewBufferString(configfiles.GetEnvironmentsJSON()))
	afero.WriteReader(fs, pagesJSON, bytes.NewBufferString(configfiles.GetPagesJSON()))
	afero.WriteReader(fs, selectorsJSON, bytes.NewBufferString(configfiles.GetSelectorsJSON()))

	if !empty {
		afero.WriteReader(fs, exampleFeature, bytes.NewBufferString(configfiles.GetExampleFeature()))
	}

	return nil
}

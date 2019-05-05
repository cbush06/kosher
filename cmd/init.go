package cmd

import (
	"bytes"
	"errors"
	"fmt"
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
	name         string
	command      *cobra.Command
	initPlatform string
	initEmpty    bool
	initForce    bool
	initFs       afero.Fs
}

var (
	validPlatforms = []string{"desktop", "web"}
)

func buildInitCommand() *initCommand {
	newCmd := &initCommand{
		name:         "init",
		command:      nil,
		initPlatform: "web",
		initFs:       &afero.OsFs{},
	}

	newCmd.command = &cobra.Command{
		Use:   "init [flags] [project name] [path to directory]",
		Short: "Initializes a new Kosher project",
		Long:  `init creates the necessary project structure with simple example tests and config files to quickly get you started.`,
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			path := ""
			if len(args) < 2 {
				if path, err = os.Getwd(); err != nil {
					return err
				}
			} else {
				if path, err = filepath.Abs(filepath.Clean(args[1])); err != nil {
					return err
				}
			}

			if platformIndex := sort.SearchStrings(validPlatforms, newCmd.initPlatform); platformIndex >= len(validPlatforms) || validPlatforms[platformIndex] != newCmd.initPlatform {
				return errors.New("Invalid platform specified [" + newCmd.initPlatform + "]. Valid options are: " + strings.Join(validPlatforms, ", "))
			}

			if err = newCmd.initProject(args[0], path, newCmd.initForce, newCmd.initEmpty); err != nil {
				return err
			}

			fmt.Printf("Project [%s] initialized...\n", args[0])
			return nil
		},
	}

	newCmd.command.Flags().StringVarP(&newCmd.initPlatform, "platform", "p", "web", "Set the platform: "+strings.Join(validPlatforms, ", "))
	newCmd.command.Flags().BoolVarP(&newCmd.initForce, "force", "f", false, "Create a project inside a non-empty directory.")
	newCmd.command.Flags().BoolVarP(&newCmd.initEmpty, "empty", "e", false, "Create an empty project (no example tests).")

	return newCmd
}

func (i *initCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(i.command)
}

// initProject initializes a new Kosher project with configuration file templates and a sample feature file
func (i *initCommand) initProject(projectName string, basepath string, force bool, empty bool) error {
	var (
		featuresDir      = filepath.Join(basepath, common.FeaturesDir)
		configDir        = filepath.Join(basepath, common.ConfigDir)
		resultsDir       = filepath.Join(basepath, common.ResultsDir)
		macrosDir        = filepath.Join(basepath, common.MacrosDir)
		resourcesDir     = filepath.Join(basepath, common.ResourcesDir)
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
		macrosDir,
		resourcesDir,
	}

	projectFiles := []string{
		environmentsJSON,
		pagesJSON,
		selectorsJSON,
		settingsJSON,
		exampleFeature,
	}

	if exists, _ := afero.Exists(i.initFs, basepath); exists {
		if isDir, _ := afero.IsDir(i.initFs, basepath); !isDir {
			return errors.New(basepath + " exists but is not a directory...aborting...")
		}

		isEmpty, _ := afero.IsEmpty(i.initFs, basepath)

		switch {
		case !isEmpty && !force:
			return errors.New(basepath + " already exists and is not empty...aborting...")

		case !isEmpty && force:
			all := append(projectStructure, projectFiles...)
			for _, path := range all {
				if pathExists, _ := afero.Exists(i.initFs, path); pathExists {
					return errors.New(path + " already exists...aborting...")
				}
			}
		}
	}

	for _, dir := range projectStructure {
		if err := i.initFs.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("Failed to create dir [%s]: %s", dir, err)
		}
	}

	if i.initPlatform == "web" {
		afero.WriteReader(i.initFs, settingsJSON, bytes.NewBufferString(configfiles.GetSettingsJSON(projectName)))
	} else if i.initPlatform == "desktop" {
		afero.WriteReader(i.initFs, settingsJSON, bytes.NewBufferString(configfiles.GetSettingsDesktopJSON(projectName)))
	}

	afero.WriteReader(i.initFs, environmentsJSON, bytes.NewBufferString(configfiles.GetEnvironmentsJSON()))
	afero.WriteReader(i.initFs, pagesJSON, bytes.NewBufferString(configfiles.GetPagesJSON()))
	afero.WriteReader(i.initFs, selectorsJSON, bytes.NewBufferString(configfiles.GetSelectorsJSON()))

	if !empty {
		afero.WriteReader(i.initFs, exampleFeature, bytes.NewBufferString(configfiles.GetExampleFeature()))
	}

	return nil
}

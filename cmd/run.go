package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/interfaces"
	"github.com/cbush06/kosher/suitecontext"

	"github.com/cbush06/godog"

	"github.com/cbush06/kosher/clients"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/report"
	"github.com/cbush06/kosher/steps/macros"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/cbush06/kosher/steps/websteps"

	"github.com/spf13/cobra"
)

type runCommand struct {
	name        string
	command     *cobra.Command
	environment string
	pathArg     string
	tags        string
	settings    *config.Settings
	fileSystem  *fs.Fs
}

func buildRunCommand() *runCommand {
	newCmd := &runCommand{
		name: "run",
	}

	newCmd.command = &cobra.Command{
		Use:   "run [flags] [path]",
		Short: "executes your tests",
		Long:  `run executes your tests. Depending on the arguments provided, it may execute all tests, a specific test, or tests in one or more subdirectories.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, arg []string) error {
			var (
				err      error
				suiteCtx *suitecontext.SuiteContext
				client   *clients.Client
			)

			// grab the path arg if specified -- this determines what feature(s) get executed
			if len(arg) < 1 {
				newCmd.pathArg, _ = os.Getwd()
				newCmd.pathArg = path.Join(newCmd.pathArg, common.FeaturesDir)
			} else {
				newCmd.pathArg = filepath.Clean(arg[0])
			}

			workingDir, _ := os.Getwd()
			if newCmd.fileSystem, err = fs.NewFs(workingDir); err != nil {
				log.Panicf("attempted to get file system but encountered error: %s", err)
				return nil
			}

			// build the settings file based on the working directory
			newCmd.settings = config.NewSettings(newCmd.fileSystem)
			newCmd.settings.Settings.BindPFlag("appVersion", cmd.Flags().Lookup("appVersion"))

			// confirm an environment is selected
			if len(newCmd.environment) < 1 {
				newCmd.environment = newCmd.settings.Settings.GetString("defaultEnvironment")
			}
			if len(newCmd.environment) < 1 {
				return errors.New(`no environment specified; either set the environment with the [run] command's [-e|--environment] flag or set the "defaultEnvironment" in the settings file`)
			}

			// verify the environment exists in the environments file and, if so, set it as the environment for this run
			if !newCmd.settings.Environments.IsSet(newCmd.environment) {
				return fmt.Errorf(`no entry found for [%s] in the environments file`, newCmd.environment)
			}
			newCmd.settings.Settings.Set("environment", newCmd.environment)

			// build the web driver client and fire it up
			if client, err = clients.NewClient(newCmd.settings); err != nil {
				return err
			}

			// Start the web driver and ensure it's shutdown later
			defer client.StopDriver()
			if err = client.StartDriver(); err != nil {
				return err
			}

			log.Printf("Web Driver server [%s] created. Serving at [%s].\n", client.DriverType, client.WebDriver.URL())

			// Create new Window
			var page interfaces.PageService
			page, err = client.WebDriver.NewPage()
			if err != nil {
				return fmt.Errorf("failed to open page: %s", err)
			}

			// Size the window
			stepUtils := steputils.NewStepUtils(newCmd.settings, page)
			if err := page.Size(stepUtils.GetMaxWindowSize()); err != nil {
				return fmt.Errorf("error encountered resizing window at startup: %s", err)
			}

			// Prepare to record results for reporting
			reportBuilder := report.NewReport(newCmd.settings)

			// Run godog if not unit testing (godog doesn't play nice with the virtual afero filesystem used during testing, so we skip godog)
			if flag.Lookup("test.v") == nil {
				godog.RunWithOptions(newCmd.settings.Settings.GetString("projectName"), func(suite *godog.Suite) {
					newCmd.buildFeatureContext(page, suite)
					suiteCtx = suitecontext.CreateSuiteContext(suite)
				}, newCmd.buildGoDogOptions(reportBuilder))
			} else {
				suiteCtx = &suitecontext.SuiteContext{}
			}

			if err := reportBuilder.Process(); err != nil {
				log.Printf("Failed to generate report: %s", err)
			}
			fmt.Printf("\nPassed: %d; Failed: %d; Pending: %d; Skipped: %d\n", suiteCtx.StepsPassed, suiteCtx.StepsFailed, suiteCtx.StepsUndefined, suiteCtx.StepsSkipped)

			return nil
		},
	}

	newCmd.command.Flags().StringVarP(&newCmd.environment, "environment", "e", "", "Set the environment.")
	newCmd.command.Flags().String("appVersion", "", "Sets the version of the application being tested for reporting purposes.")
	newCmd.command.Flags().StringVarP(&newCmd.tags, "tags", "t", "", "Filter features, scenarios, scenario outlines, and examples by tags.")

	return newCmd
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(r.command)
}

func (r *runCommand) buildFeatureContext(page interfaces.PageService, suite *godog.Suite) {
	// Load primary steps based on platform
	switch r.settings.Settings.Get("platform") {
	case "desktop":
		log.Fatal("desktop is not implemented")
	case "web":
		websteps.BuildGoDogSuite(r.settings, page, suite)
	}

	// Load macro steps
	if macros, err := macros.BuildMacros(r.settings.FileSystem); err != nil {
		log.Fatalf("error encountered while parsing macros: %s", err)
	} else {
		for _, m := range macros {
			nextMacro := m
			suite.Step("^"+nextMacro.Step+"$", func() godog.Steps { return *nextMacro.Substeps })
		}
	}
}

func (r *runCommand) buildGoDogOptions(reportBuilder report.Report) godog.Options {
	featuresPath, _ := filepath.Abs(r.pathArg)

	// Convert kosher format to GoDog format
	var reportFormat string
	switch r.settings.Settings.GetString("reportFormat") {
	case "html", "bootstrap", "simple":
		reportFormat = "cucumber"
	default:
		reportFormat = r.settings.Settings.GetString("reportFormat")
	}

	return godog.Options{
		Format:        reportFormat,
		Paths:         []string{featuresPath},
		Tags:          r.tags,
		StopOnFailure: r.settings.Settings.GetBool("quitOnFail"),
		Strict:        r.settings.Settings.GetBool("quitOnFail"),
		Output:        reportBuilder,
	}
}

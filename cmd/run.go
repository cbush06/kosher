package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/sclevine/agouti"

	"github.com/cbush06/kosher/clients"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/report"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/cbush06/kosher/steps/websteps"

	"github.com/spf13/cobra"
)

type runCommand struct {
	name    string
	command *cobra.Command
}

var (
	err         error
	environment string
	pathArg     string
	fileSys     *fs.Fs
	settings    *config.Settings
)

var cmdRun = &runCommand{
	name: "run",
	command: &cobra.Command{
		Use:   "run [flags] [path]",
		Short: "executes your tests",
		Long:  `run executes your tests. Depending on the arguments provided, it may execute all tests, a specific test, or tests in one or more subdirectories.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, arg []string) error {
			// grab the path arg if specified -- this determines what feature(s) get executed
			if len(arg) < 1 {
				pathArg, _ = os.Getwd()
			} else {
				pathArg = filepath.Clean(arg[0])
			}

			// determine where the executable was called from
			workingDir, _ := os.Getwd()
			if fileSys, err = fs.NewFs(workingDir); err != nil {
				log.Fatal(err)
			}

			// build the settings file based on the working directory
			settings = config.NewSettings(fileSys)
			settings.Settings.BindPFlag("appVersion", cmd.Flags().Lookup("appVersion"))

			// confirm an environment is selected
			if len(environment) < 1 {
				environment = settings.Settings.GetString("defaultEnvironment")
			}
			if len(environment) < 1 {
				log.Fatal(`No environment specified. Either set the environment with the [run] command's [-e|--environment] flag or set the "defaultEnvironment" in the settings file.`)
			}

			// verify the environment exists in the environments file and, if so, set it as the environment for this run
			if !settings.Environments.IsSet(environment) {
				log.Fatalf(`No entry found for [%s] in the environments file.`, environment)
			} else {
				settings.Settings.Set("environment", environment)
			}

			client, err := clients.NewClient(settings, fileSys)
			if err != nil {
				log.Fatal(err)
			} else {
				defer client.StopDriver()
				client.StartDriver()

				log.Printf("Web Driver server [%s] created. Serving at [%s].\n", client.DriverType, client.WebDriver.URL())

				// Create new Window
				page, err := client.WebDriver.NewPage()
				if err != nil {
					log.Fatalf("failed to open page: %s", err)
				}

				// Size the window
				stepUtils := steputils.NewStepUtils(settings, page)
				if err := page.Size(stepUtils.GetMaxWindowSize()); err != nil {
					return fmt.Errorf("error encountered resizing window at startup: %s", err)
				}

				reportBuilder := report.NewReport(settings, fileSys)
				godog.RunWithOptions(settings.Settings.GetString("projectName"), func(suite *godog.Suite) {
					buildFeatureContext(settings, page, suite)
				}, buildGoDogOptions(settings, reportBuilder))

				if err := reportBuilder.Process(); err != nil {
					log.Printf("Failed to generate report: %s", err)
				}
			}

			return nil
		},
	},
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
	r.command.Flags().StringVarP(&environment, "environment", "e", "", "Set the environment")
	r.command.Flags().String("appVersion", "", "Sets the version of the application being tested for reporting purposes.")
	cmd.AddCommand(r.command)
}

func buildFeatureContext(settings *config.Settings, page *agouti.Page, suite *godog.Suite) {
	switch settings.Settings.Get("platform") {
	case "desktop":
		log.Fatal("desktop is not implemented")
	case "web":
		websteps.BuildGoDogSuite(settings, page, suite)
	}
}

func buildGoDogOptions(settings *config.Settings, reportBuilder report.Report) godog.Options {
	featuresPath, _ := filepath.Abs(pathArg)
	// featuresPath := pathArg // TODO: remove this and use above line once GoDog fixes their windows issue

	// Convert kosher format to GoDog format
	var reportFormat string
	switch settings.Settings.GetString("reportFormat") {
	case "html", "bootstrap", "simple":
		reportFormat = "cucumber"
	default:
		reportFormat = settings.Settings.GetString("reportFormat")
	}

	return godog.Options{
		Format:        reportFormat,
		Paths:         []string{featuresPath},
		StopOnFailure: settings.Settings.GetBool("quitOnFail"),
		Strict:        settings.Settings.GetBool("quitOnFail"),
		Output:        reportBuilder,
	}
}

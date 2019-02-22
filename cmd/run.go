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
		Use:   "run",
		Short: "executes your tests",
		Long:  `run executes your tests. Depending on the arguments provided, it may execute all tests, a specific test, or tests in one or more subdirectories.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, arg []string) error {
			if len(arg) < 1 {
				pathArg, _ = os.Getwd()
			} else {
				pathArg = filepath.Clean(arg[0])
			}

			fmt.Printf("App Version: %s\n", settings.Settings.GetString("appversion"))

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

				fmt.Printf("Web Driver server [%s] created. Serving at [%s].\n", client.DriverType, client.WebDriver.URL())

				page, err := client.WebDriver.NewPage(agouti.Browser("chrome"))
				if err != nil {
					log.Fatalf("failed to open page: %s", err)
				}

				godog.RunWithOptions(settings.Settings.GetString("projectName"), func(suite *godog.Suite) {
					buildFeatureContext(settings, page, suite)
				}, buildGoDogOptions(settings))
			}

			return nil
		},
	},
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
	workingDir, _ := os.Getwd()
	if fileSys, err = fs.NewFs(workingDir); err != nil {
		log.Fatal(err)
	}

	settings = config.NewSettings(fileSys)

	r.command.Flags().StringVarP(&environment, "environment", "e", "", "Set the environment")
	r.command.Flags().String("appversion", "", "Sets the version of the application being tested for reporting purposes.")
	settings.Settings.BindPFlag("appversion", r.command.Flags().Lookup("appversion"))
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

func buildGoDogOptions(settings *config.Settings) godog.Options {
	featuresPath, _ := filepath.Abs(pathArg)

	fmt.Println(featuresPath)

	return godog.Options{
		Format:        settings.Settings.GetString("reportFormat"),
		Paths:         []string{featuresPath},
		StopOnFailure: settings.Settings.GetBool("quitOnFail"),
		Strict:        settings.Settings.GetBool("quitOnFail"),
	}
}

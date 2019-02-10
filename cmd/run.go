package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cbush06/kosher/common"

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
			client, err := clients.NewClient(settings, fs)
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
				}, buildGoDogOptions(settings, fs))
			}

			return nil
		},
	},
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
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

func buildGoDogOptions(settings *config.Settings, fs *fs.Fs) godog.Options {
	featuresPath, err := fs.ProjectDir.RealPath(common.FeaturesDir)
	if err != nil {
		log.Fatalln("Failed to get path to features directory")
	}

	return godog.Options{
		Format: settings.Settings.GetString("reportFormat"),
		Paths:  []string{featuresPath},
	}
}

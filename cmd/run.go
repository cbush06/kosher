package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sclevine/agouti"

	"github.com/cbush06/kosher/clients"
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
			client, err := clients.NewClient(settings, fs)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				defer client.StopDriver()
				client.StartDriver()

				fmt.Printf("Web Driver server [%s] created. Serving at [%s].\n", client.DriverType, client.WebDriver.URL())

				page, err := client.WebDriver.NewPage(agouti.Browser("chrome"))
				if err != nil {
					log.Fatalf("failed to open page: %s", err.Error())
				}
				if err := page.Navigate("http://www.drudgereport.com"); err != nil {
					log.Fatalf("failed to navigate: %s", err.Error())
				}
				time.Sleep(3 * time.Second)
				title, err := page.Title()
				if err != nil {
					log.Fatalf("failed to get title: %s", err.Error())
				}
				fmt.Println(title)
			}

			return nil
		},
	},
}

func (r *runCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(r.command)
}

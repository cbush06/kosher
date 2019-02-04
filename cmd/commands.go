package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = cmdKosher.command

// Init prepares the CLI interpreter and executes it
func Init() {
	cobra.OnInitialize(initConfig)
	cmdVersion.registerWith(rootCmd)
	cmdInit.registerWith(rootCmd)
	cmdRun.registerWith(rootCmd)
	rootCmd.Execute()
}

func initConfig() {
	fmt.Println("Initializing config")
}

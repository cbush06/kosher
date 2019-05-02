package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = cmdKosher.command

// Init prepares the CLI interpreter and executes it
func Init() {
	cobra.OnInitialize(configureEnv)

	// Register main commands and set flags of commands
	cmdVersion.registerWith(rootCmd)

	buildInitCommand().registerWith(rootCmd)

	cmdRun.registerWith(rootCmd)
	cmdRun.setFlags()

	cmdSend.registerWith(rootCmd)

	// Register `send` commands
	cmdJira.registerWith(cmdSend.command)
	cmdJira.setFlags()

	rootCmd.Execute()
}

func configureEnv() {
	var pathEnv string
	if pathEnv = os.Getenv("PATH"); len(pathEnv) < 1 {
		log.Fatalf("Unable to retrieve PATH environment variable value")
	}

	libsDir, _ := exec.LookPath(os.Args[0])
	libsDir = filepath.Join(filepath.Dir(libsDir), "libs")

	switch runtime.GOOS {
	case "linux":
		os.Setenv("PATH", pathEnv+":"+libsDir)
	case "windows":
		os.Setenv("PATH", pathEnv+";"+libsDir)
	default:
		log.Fatalf("Unsupported operating system detected [%s]\n", runtime.GOOS)
	}

}

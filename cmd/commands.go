package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = cmdKosher.command

// Init prepares the CLI interpreter and executes it
func Init() {
	cobra.OnInitialize(configureEnv)
	cmdVersion.registerWith(rootCmd)
	cmdInit.registerWith(rootCmd)
	cmdRun.registerWith(rootCmd)
	rootCmd.Execute()
}

func configureEnv() {
	var pathEnv string
	if pathEnv = os.Getenv("PATH"); len(pathEnv) < 1 {
		log.Fatalf("Unable to retrieve PATH environment variable value")
	}

	libsDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	libsDir = filepath.Join(libsDir, "libs")

	switch runtime.GOOS {
	case "linux":
		os.Setenv("PATH", pathEnv+":"+libsDir)
	case "windows":
		os.Setenv("PATH", pathEnv+";"+libsDir)
	default:
		log.Fatalf("Unsupported operating system detected [%s]\n", runtime.GOOS)
	}

}

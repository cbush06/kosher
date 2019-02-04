package kosherfs

import (
	"github.com/kosher/config"
	"github.com/spf13/afero"
)

// Os points to an Os Afero file system.
var Os = &afero.OsFs{}

// Fs abstracts away the file system from the actual OS, and allows the file system to be mocked for testing.
type Fs struct {
	// Main working directory for Kosher
	Project afero.Fs

	// Read-only view of working directory for Kosher
	WorkingDir *afero.BasePathFs
}

// NewDefault creates a new Fs with the OS file system as the Project directory
func NewDefault(cfg config.Provider) *Fs {
	fs := &afero.OsFs{}
	return newFs(fs, cfg)
}

func newFs(base afero.Fs, cfg config.Provider) *Fs {
	return &Fs{
		Project:    base,
		WorkingDir: getWorkingDirFs(base, cfg),
	}
}

func getWorkingDirFs(base afero.Fs, cfg config.Provider) *afero.BasePathFs {
	workingDir := cfg.GetString("workingDir")

	if workingDir != "" {
		return afero.NewBasePathFs(afero.NewReadOnlyFs(base), workingDir).(*afero.BasePathFs)
	}

	return nil
}

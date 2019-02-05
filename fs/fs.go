package fs

import (
	"github.com/cbush06/kosher/common"
	"github.com/spf13/afero"
)

// Os points to an Os Afero file system.
var Os = &afero.OsFs{}

// Fs abstracts away the file system from the actual OS, and allows the file system to be mocked for testing.
type Fs struct {
	// Read-only view of working directory for Kosher (this is the root dir)
	WorkingDir *afero.BasePathFs

	// Config directory
	ConfigDir *afero.BasePathFs

	// Features directory
	FeaturesDir *afero.BasePathFs
}

// NewFs creates a new Fs with the OS file system as the Project directory
func NewFs(workingDirPath string) *Fs {
	fs := &afero.OsFs{}
	return newFs(fs, workingDirPath)
}

func newFs(base afero.Fs, workingDirPath string) *Fs {
	var workingDir = getWorkingDirFs(base, workingDirPath)

	return &Fs{
		WorkingDir:  workingDir,
		ConfigDir:   afero.NewBasePathFs(afero.NewReadOnlyFs(workingDir), common.ConfigDir).(*afero.BasePathFs),
		FeaturesDir: afero.NewBasePathFs(afero.NewReadOnlyFs(workingDir), common.FeaturesDir).(*afero.BasePathFs),
	}
}

func getWorkingDirFs(base afero.Fs, workingDirPath string) *afero.BasePathFs {
	if workingDirPath != "" {
		return afero.NewBasePathFs(afero.NewReadOnlyFs(base), workingDirPath).(*afero.BasePathFs)
	}

	return nil
}

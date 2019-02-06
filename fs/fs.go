package fs

import (
	"errors"
	"path/filepath"

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

	// Results directory
	ResultsDir *afero.BasePathFs
}

// NewFs creates a new Fs with the OS file system as the Project directory
func NewFs(workingDirPath string) (*Fs, error) {
	fs := &afero.OsFs{}
	return newFs(fs, workingDirPath)
}

func newFs(base afero.Fs, workingDirPath string) (*Fs, error) {
	var (
		workingDir  *afero.BasePathFs
		configDir   *afero.BasePathFs
		featuresDir *afero.BasePathFs
		resultsDir  *afero.BasePathFs
		err         error
	)

	workingDirPath, err = filepath.Abs(workingDirPath)
	if workingDir, err = getWorkingDirFs(base, workingDirPath); err != nil {
		return nil, err
	}

	configDirPath, _ := filepath.Abs(filepath.Join(workingDirPath, common.ConfigDir))
	if exists, _ := afero.DirExists(base, configDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + configDirPath)
	}
	configDir = afero.NewBasePathFs(workingDir, common.ConfigDir).(*afero.BasePathFs)

	featuresDirPath, _ := filepath.Abs(filepath.Join(workingDirPath, common.FeaturesDir))
	if exists, _ := afero.DirExists(base, featuresDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + featuresDirPath)
	}
	featuresDir = afero.NewBasePathFs(workingDir, common.FeaturesDir).(*afero.BasePathFs)

	resultsDirPath, _ := filepath.Abs(filepath.Join(workingDirPath, common.ResultsDir))
	if exists, _ := afero.DirExists(base, resultsDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + resultsDirPath)
	}
	resultsDir = afero.NewBasePathFs(workingDir, common.ResultsDir).(*afero.BasePathFs)

	return &Fs{
		WorkingDir:  workingDir,
		ConfigDir:   configDir,
		FeaturesDir: featuresDir,
		ResultsDir:  resultsDir,
	}, nil
}

func getWorkingDirFs(base afero.Fs, workingDirPath string) (*afero.BasePathFs, error) {
	if workingDirPath != "" {
		if exists, _ := afero.DirExists(base, workingDirPath); !exists {
			return nil, errors.New("Directory does not exist: " + workingDirPath)
		}
		return afero.NewBasePathFs(afero.NewReadOnlyFs(base), workingDirPath).(*afero.BasePathFs), nil
	}
	return nil, errors.New("No working directory path provided")
}

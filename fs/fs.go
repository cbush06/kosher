package fs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/cbush06/kosher/common"
	"github.com/spf13/afero"
)

// Fs abstracts away the file system from the actual OS, and allows the file system to be mocked for testing.
type Fs struct {
	// Read-only view of working directory for Kosher (where the executable is)
	WorkingDir *afero.BasePathFs

	// Read-only view of project directory for Kosher
	ProjectDir *afero.BasePathFs

	// Config directory
	ConfigDir *afero.BasePathFs

	// Features directory
	FeaturesDir *afero.BasePathFs

	// Results directory
	ResultsDir *afero.BasePathFs

	// Macros directory
	MacrosDir *afero.BasePathFs
}

// NewFs creates a new Fs with the OS file system as the Project directory
func NewFs(projectDirPath string) (*Fs, error) {
	fs := afero.NewOsFs()
	return newFs(fs, projectDirPath)
}

func newFs(base afero.Fs, projectDirPath string) (*Fs, error) {
	var (
		workingDirPath string
		workingDir     *afero.BasePathFs
		projectDir     *afero.BasePathFs
		configDir      *afero.BasePathFs
		featuresDir    *afero.BasePathFs
		resultsDir     *afero.BasePathFs
		macrosDir      *afero.BasePathFs
		err            error
	)

	if workingDirPath, err = os.Getwd(); err != nil {
		return nil, errors.New("Unable to determine working directory of executable")
	}
	if workingDir, err = getReadonlyBasepathFs(base, workingDirPath); err != nil {
		return nil, err
	}

	projectDirPath, err = filepath.Abs(projectDirPath)
	if projectDir, err = getReadonlyBasepathFs(base, projectDirPath); err != nil {
		return nil, err
	}

	configDirPath, _ := filepath.Abs(filepath.Join(projectDirPath, common.ConfigDir))
	if exists, _ := afero.DirExists(base, configDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + configDirPath)
	}
	configDir = afero.NewBasePathFs(projectDir, common.ConfigDir).(*afero.BasePathFs)

	featuresDirPath, _ := filepath.Abs(filepath.Join(projectDirPath, common.FeaturesDir))
	if exists, _ := afero.DirExists(base, featuresDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + featuresDirPath)
	}
	featuresDir = afero.NewBasePathFs(projectDir, common.FeaturesDir).(*afero.BasePathFs)

	resultsDirPath, _ := filepath.Abs(filepath.Join(workingDirPath, common.ResultsDir))
	if exists, _ := afero.DirExists(base, resultsDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + resultsDirPath)
	}
	resultsDir = afero.NewBasePathFs(base, common.ResultsDir).(*afero.BasePathFs)

	macrosDirPath, _ := filepath.Abs(filepath.Join(projectDirPath, common.MacrosDir))
	if exists, _ := afero.DirExists(base, macrosDirPath); !exists {
		return nil, errors.New("Directory does not exist: " + macrosDirPath)
	}
	macrosDir = afero.NewBasePathFs(projectDir, common.MacrosDir).(*afero.BasePathFs)

	return &Fs{
		WorkingDir:  workingDir,
		ProjectDir:  projectDir,
		ConfigDir:   configDir,
		FeaturesDir: featuresDir,
		ResultsDir:  resultsDir,
		MacrosDir:   macrosDir,
	}, nil
}

func getReadonlyBasepathFs(base afero.Fs, dirPath string) (*afero.BasePathFs, error) {
	if dirPath != "" {
		if exists, _ := afero.DirExists(base, dirPath); !exists {
			return nil, errors.New("Directory does not exist: " + dirPath)
		}
		return afero.NewBasePathFs(afero.NewReadOnlyFs(base), dirPath).(*afero.BasePathFs), nil
	}
	return nil, errors.New("No directory path provided")
}

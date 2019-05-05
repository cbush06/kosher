package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cbush06/kosher/common"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestInitCommandArgValidation(t *testing.T) {
	t.Run("0 args", func(t *testing.T) {
		var cmdInit = buildInitCommand()
		cmdInit.command.SetArgs([]string{})
		result := cmdInit.command.Execute()
		assert.NotNil(t, result, "expected init command to fail due to no args received when [1,2] required, but did not fail")
	})

	t.Run("3 args", func(t *testing.T) {
		var cmdInit = buildInitCommand()
		cmdInit.command.SetArgs([]string{"projname", "projpath", "bogusarg"})
		result := cmdInit.command.Execute()
		assert.NotNil(t, result, "expected init command to fail due to 3 args received when [1,2] required, but did not fail")
	})
}

func TestInitCommandPlatformValidation(t *testing.T) {
	t.Run("test platform flag", func(t *testing.T) {
		for _, nextPlatform := range validPlatforms {
			t.Run("valid platform ("+nextPlatform+")", func(t *testing.T) {
				var cmdInit = buildInitCommand()
				cmdInit.initFs = afero.NewMemMapFs()
				cmdInit.command.SetArgs([]string{"projname", "-p", "web"})
				result := cmdInit.command.Execute()
				assert.Nilf(t, result, "expected init command to succeed, but did not %s", result)
			})
		}

		t.Run("invalid platform", func(t *testing.T) {
			var cmdInit = buildInitCommand()
			cmdInit.initFs = afero.NewMemMapFs()
			cmdInit.command.SetArgs([]string{"projname", "-p", "bogus"})
			result := cmdInit.command.Execute()
			assert.NotNil(t, result, "expected init command to fail, but did not")
		})
	})
}

func TestInitCommandDirectoryCreation(t *testing.T) {
	arg1 := "projname"
	basePath, _ := os.Getwd()
	arg2 := "somedir/someotherdir"
	arg2Path, _ := filepath.Abs(arg2)

	t.Run("working_directory", func(t *testing.T) {
		var cmdInit = buildInitCommand()
		cmdInit.initFs = afero.NewMemMapFs()

		cmdInit.command.SetArgs([]string{arg1})
		result := cmdInit.command.Execute()

		assert.Nilf(t, result, "expected init command to run without errors, but received: %s", result)

		var dirExists bool

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.FeaturesDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.FeaturesDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ConfigDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ConfigDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ResultsDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ResultsDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.MacrosDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.MacrosDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ResourcesDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ResourcesDir)

		var fileExists bool

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ConfigDir, common.EnvironmentsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.EnvironmentsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ConfigDir, common.PagesFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.PagesFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ConfigDir, common.SelectorsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.SelectorsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.ConfigDir, common.SettingsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.SettingsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(basePath, common.FeaturesDir, common.ExampleFeatureFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.ExampleFeatureFile)
	})

	t.Run("specified_directory", func(t *testing.T) {
		var cmdInit = buildInitCommand()
		cmdInit.initFs = afero.NewMemMapFs()

		cmdInit.command.SetArgs([]string{arg1, arg2})
		result := cmdInit.command.Execute()

		assert.Nilf(t, result, "expected init command to run without errors, but received: %s", result)

		var dirExists bool

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.FeaturesDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.FeaturesDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ConfigDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ConfigDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ResultsDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ResultsDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.MacrosDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.MacrosDir)

		dirExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ResourcesDir))
		assert.Truef(t, dirExists, "expected /%s directory but none found", common.ResourcesDir)

		var fileExists bool

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ConfigDir, common.EnvironmentsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.EnvironmentsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ConfigDir, common.PagesFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.PagesFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ConfigDir, common.SelectorsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.SelectorsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.ConfigDir, common.SettingsFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.SettingsFile)

		fileExists, _ = afero.Exists(cmdInit.initFs, filepath.Join(arg2Path, common.FeaturesDir, common.ExampleFeatureFile))
		assert.Truef(t, fileExists, "expected file %s but none found", common.ExampleFeatureFile)
	})
}

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/cbush06/kosher/fs"
	"github.com/sclevine/agouti"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/interfaces"
	"github.com/spf13/afero"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandArgValidation(t *testing.T) {
	// If we're unit testing, construct a sample project in memory
	fs.MockFs = afero.NewMemMapFs()
	common.BuildTestProject(fs.MockFs)

	t.Run("0 args", func(t *testing.T) {
		var cmdRun = buildRunCommand()
		cmdRun.command.SetArgs([]string{})

		// add dummy config files
		afero.WriteReader(cmdRun.fileSystem.ConfigDir, common.SettingsFile, bytes.NewBufferString(`{"driver":"mock"}`))
		afero.WriteReader(cmdRun.fileSystem.ConfigDir, common.EnvironmentsFile, bytes.NewBufferString(`{"test":"https://www.google.com"}`))

		// Prepare the MockPage for assertions
		page := new(interfaces.MockPage)
		page.On("Size", 2000, 980).Return(nil).Once()
		page.On("Session").Return(new(interfaces.MockSession)).Once()

		// Prepare the MockDriver for assertions
		driver := interfaces.UnitTestingMockDriver
		driver.On("Start").Return(nil).Once()
		driver.On("Stop").Return(nil).Once()
		driver.On("NewPage", []agouti.Option(nil)).Return(page, nil).Once()
		driver.On("URL").Return("").Once()

		// Test our assertions
		result := cmdRun.command.Execute()
		assert.Nilf(t, result, "expected run command to succeed, but received error: %s", result)

		path, _ := os.Getwd()
		assert.Equal(t, filepath.Join(path, common.FeaturesDir), cmdRun.pathArg, "invalid features path set")
		assert.NotNil(t, cmdRun.environment, "environment should be set but is not")

		page.AssertExpectations(t)
	})
}

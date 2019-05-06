package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/fs"
	"github.com/cbush06/kosher/integrations"
	"github.com/spf13/afero"
)

func TestJiraCommandDefaultPath(t *testing.T) {
	// If we're unit testing, consruct a sample project in memory
	fs.MockFs = afero.NewMemMapFs()
	afero.WriteReader(fs.MockFs, filepath.Join(common.ResultsDir, common.ResultsJSONFile), bytes.NewBufferString("{}"))
	integrations.MockSendable.On("Send", mock.Anything, mock.Anything).Return(nil).Once()
	common.BuildTestProject(fs.MockFs)

	var cmdJira = buildJiraCommand()
	cmdJira.command.Execute()

	integrations.MockSendable.AssertNumberOfCalls(t, "Send", 1)
}

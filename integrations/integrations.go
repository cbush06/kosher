package integrations

import (
	"flag"
	"fmt"
	"os"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/integrations/jira"
	"github.com/cbush06/kosher/integrations/mocksendable"
	"github.com/cbush06/kosher/report"
	"github.com/spf13/afero"
)

const (
	// Jira specifies a Jira system to `SendTo` command
	Jira = iota

	// Mock specifies a mock Sendable
	Mock
)

// MockSendable is a mock sendable for unit testing
var MockSendable = new(mocksendable.MockSendable)

// SendTo transmits the results stored in the `results.json` file to the specified `system`. The `system` parameter should be one of the constants specified in this file.
func SendTo(system int, s *config.Settings) error {
	var (
		cukeReport  = report.NewCucumberReport(s)
		fileExists  bool
		jsonFile    afero.File
		jsonResults []byte
		sendable    Sendable
		err         error
	)

	// Ensure the results.json file exists
	if fileExists, err = afero.Exists(s.FileSystem.ResultsDir, common.ResultsJSONFile); err != nil {
		return fmt.Errorf("error encountered while checking for [%s] file: %s", common.ResultsJSONFile, err)
	} else if !fileExists {
		return fmt.Errorf("no [%s] file found; you may need to run your kosher tests with [reportFormat] in [%s] set to cucumber, simple, or bootstrap", common.ResultsJSONFile, common.SettingsFile)
	}

	// Open the JSON results file
	if jsonFile, err = s.FileSystem.ResultsDir.OpenFile(common.ResultsJSONFile, os.O_RDONLY, 0744); err != nil {
		return fmt.Errorf("Error encountered while opening results file [%s]: %s", common.ResultsJSONFile, err)
	}

	// Read the file in and ensure it was read in full
	if jsonResults, err = afero.ReadAll(jsonFile); err != nil {
		return fmt.Errorf("Error reading results file [%s]: %s", common.ResultsJSONFile, err)
	}

	// Unmarshall the file into a CucumberReport
	cukeReport.UnmarshallJSON(jsonResults)

	// Call the appropriate integration (use mock for unit testing)
	if flag.Lookup("test.v") != nil {
		sendable = MockSendable
	} else {
		switch system {
		case Jira:
			sendable = new(jira.Jira)
		}
	}

	return sendable.Send(s, &cukeReport)
}

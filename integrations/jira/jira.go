package jira

import (
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/report"
)

// Jira represents a connection to a Jira server.
type Jira struct {
}

// Send connects to the configured Jira server, retrieves the user's credentials
// via CLI, and creates new issues for failed tests in the CucumberReport.
func (j *Jira) Send(s *config.Settings, r *report.CucumberReport) error {
	return nil
}

package integrations

import (
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/report"
)

// Sendable abstracts the ability of an object to send a CucumberReport to an interfacing system.
type Sendable interface {
	Send(settings *config.Settings, cukeReport *report.CucumberReport) error
}

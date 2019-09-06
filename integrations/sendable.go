package integrations

import (
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/formats"
)

// Sendable abstracts the ability of an object to send a CucumberReport to an interfacing system.
type Sendable interface {
	Send(*config.Settings, []formats.CukeFeature) error
}

package mocksendable

import (
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/report"
	"github.com/stretchr/testify/mock"
)

// MockSendable mocks a sendable for unit testing
type MockSendable struct {
	mock.Mock
}

// Send mocks the Sendable interface's SendTo method
func (m *MockSendable) Send(settings *config.Settings, cukeReport *report.CucumberReport) error {
	args := m.Called()
	return args.Error(0)
}

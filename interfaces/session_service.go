package interfaces

import (
	"github.com/sclevine/agouti/api"
	"github.com/stretchr/testify/mock"
)

// SessionService allows mocking agouti.api.Session
type SessionService interface {
	GetActiveElement() (*api.Element, error)
	GetWindows() ([]*api.Window, error)
	SetWindow(window *api.Window) error
}

// MockSession serves as a mock for unit testing
type MockSession struct {
	mock.Mock
}

// GetActiveElement is a mock method
func (m *MockSession) GetActiveElement() (*api.Element, error) {
	args := m.Called()
	return args.Get(0).(*api.Element), args.Error(1)
}

// GetWindows is a mock method
func (m *MockSession) GetWindows() ([]*api.Window, error) {
	args := m.Called()
	return args.Get(0).([]*api.Window), args.Error(1)
}

// SetWindow is a mock method
func (m *MockSession) SetWindow(window *api.Window) error {
	args := m.Called()
	return args.Error(0)
}

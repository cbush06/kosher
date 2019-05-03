package interfaces

import (
	"github.com/sclevine/agouti/api"
	"github.com/stretchr/testify/mock"
)

// PageService allows mocking of agouti.Page
type PageService interface {
	SelectorService
	Size(width, height int) error
	Screenshot(filename string) error
	WindowCount() (int, error)
	Session() *api.Session
	Refresh() error
	RunScript(body string, arguments map[string]interface{}, result interface{}) error
	Navigate(url string) error
	SwitchToRootFrame() error
	ConfirmPopup() error
	CancelPopup() error
	EnterPopupText(text string) error
	PopupText() (string, error)
	URL() (string, error)
}

// MockPage is a mock of agouti.Page
type MockPage struct {
	MockSelectable
	mock.Mock
}

// Size is a mock method
func (m *MockPage) Size(width, height int) error {
	args := m.Called(width, height)
	return args.Error(0)
}

// Screenshot is a mock method
func (m *MockPage) Screenshot(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

// WindowCount is a mock method
func (m *MockPage) WindowCount() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

// Session is a mock method
func (m *MockPage) Session() *api.Session {
	args := m.Called()
	return args.Get(0).(*api.Session)
}

// Refresh is a mock method
func (m *MockPage) Refresh() error {
	args := m.Called()
	return args.Error(0)
}

// RunScript is a mock method
func (m *MockPage) RunScript(body string, arguments map[string]interface{}, result interface{}) error {
	args := m.Called(body, arguments, result)
	return args.Error(0)
}

// Navigate is a mock method
func (m *MockPage) Navigate(url string) error {
	args := m.Called(url)
	return args.Error(0)
}

// SwitchToRootFrame is a mock method
func (m *MockPage) SwitchToRootFrame() error {
	args := m.Called()
	return args.Error(0)
}

// ConfirmPopup is a mock method
func (m *MockPage) ConfirmPopup() error {
	args := m.Called()
	return args.Error(0)
}

// CancelPopup is a mock method
func (m *MockPage) CancelPopup() error {
	args := m.Called()
	return args.Error(0)
}

// EnterPopupText is a mock method
func (m *MockPage) EnterPopupText(text string) error {
	args := m.Called(text)
	return args.Error(0)
}

// PopupText is a mock method
func (m *MockPage) PopupText() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// URL is a mock method
func (m *MockPage) URL() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

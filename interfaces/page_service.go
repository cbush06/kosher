package interfaces

import (
	"github.com/sclevine/agouti"
	"github.com/stretchr/testify/mock"
)

// PageService allows mocking of agouti.Page
type PageService interface {
	SelectorService
	Size(width, height int) error
	Screenshot(filename string) error
	WindowCount() (int, error)
	Session() SessionService
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

// NewPageServiceFromAgoutiPage either wraps an agouti.Page (if one is provided) or a MockPage (if nil is provided) in a PageService
func NewPageServiceFromAgoutiPage(page *agouti.Page) PageService {
	return &AgoutiPage{
		SelectorService: NewSelectorServiceFromAgoutiPage(page),
		page:            page,
	}
}

// AgoutiPage wraps an agouti.Page as an adapter
type AgoutiPage struct {
	SelectorService
	page *agouti.Page
}

// Size is a mock method
func (m *AgoutiPage) Size(width, height int) error {
	return m.page.Size(width, height)
}

// Screenshot is a mock method
func (m *AgoutiPage) Screenshot(filename string) error {
	return m.page.Screenshot(filename)
}

// WindowCount is a mock method
func (m *AgoutiPage) WindowCount() (int, error) {
	return m.page.WindowCount()
}

// Session is a mock method
func (m *AgoutiPage) Session() SessionService {
	return m.page.Session()
}

// Refresh is a mock method
func (m *AgoutiPage) Refresh() error {
	return m.page.Refresh()
}

// RunScript is a mock method
func (m *AgoutiPage) RunScript(body string, arguments map[string]interface{}, result interface{}) error {
	return m.page.RunScript(body, arguments, result)
}

// Navigate is a mock method
func (m *AgoutiPage) Navigate(url string) error {
	return m.page.Navigate(url)
}

// SwitchToRootFrame is a mock method
func (m *AgoutiPage) SwitchToRootFrame() error {
	return m.page.SwitchToRootFrame()
}

// ConfirmPopup is a mock method
func (m *AgoutiPage) ConfirmPopup() error {
	return m.page.ConfirmPopup()
}

// CancelPopup is a mock method
func (m *AgoutiPage) CancelPopup() error {
	return m.page.CancelPopup()
}

// EnterPopupText is a mock method
func (m *AgoutiPage) EnterPopupText(text string) error {
	return m.page.EnterPopupText(text)
}

// PopupText is a mock method
func (m *AgoutiPage) PopupText() (string, error) {
	return m.page.PopupText()
}

// URL is a mock method
func (m *AgoutiPage) URL() (string, error) {
	return m.page.URL()
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
func (m *MockPage) Session() SessionService {
	args := m.Called()
	return args.Get(0).(SessionService)
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

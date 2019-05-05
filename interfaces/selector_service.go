package interfaces

import (
	"github.com/sclevine/agouti"
	"github.com/stretchr/testify/mock"
)

// SelectorService allows mocking agouti selectables
type SelectorService interface {
	All(selector string) *agouti.MultiSelection
	AllByXPath(selector string) *agouti.MultiSelection
	AllByLabel(text string) *agouti.MultiSelection
	AllByButton(text string) *agouti.MultiSelection
	AllByName(name string) *agouti.MultiSelection
	AllByLink(text string) *agouti.MultiSelection
	AllByID(text string) *agouti.MultiSelection
	AllByClass(text string) *agouti.MultiSelection
	FirstByLink(text string) *agouti.Selection
	First(selector string) *agouti.Selection
}

// NewSelectorServiceFromAgoutiPage either wraps an agouti.Page (if one is provided) or a MockSelector (if nil is provided) in a SelectorService
func NewSelectorServiceFromAgoutiPage(page *agouti.Page) SelectorService {
	return &AgoutiSelector{
		page: page,
	}
}

// AgoutiSelector wraps an *agouti.Page to adapt its selectable methods
type AgoutiSelector struct {
	page *agouti.Page
}

// All is a mock method
func (m *AgoutiSelector) All(selector string) *agouti.MultiSelection {
	return m.page.All(selector)
}

// AllByXPath is a mock method
func (m *AgoutiSelector) AllByXPath(selector string) *agouti.MultiSelection {
	return m.page.AllByXPath(selector)
}

// AllByLabel is a mock method
func (m *AgoutiSelector) AllByLabel(text string) *agouti.MultiSelection {
	return m.page.AllByLabel(text)
}

// AllByButton is a mock method
func (m *AgoutiSelector) AllByButton(text string) *agouti.MultiSelection {
	return m.page.AllByButton(text)
}

// AllByName is a mock method
func (m *AgoutiSelector) AllByName(name string) *agouti.MultiSelection {
	return m.page.AllByName(name)
}

// AllByLink is a mock method
func (m *AgoutiSelector) AllByLink(text string) *agouti.MultiSelection {
	return m.page.AllByLink(text)
}

// AllByID is a mock method
func (m *AgoutiSelector) AllByID(text string) *agouti.MultiSelection {
	return m.page.AllByID(text)
}

// AllByClass is a mock method
func (m *AgoutiSelector) AllByClass(text string) *agouti.MultiSelection {
	return m.page.AllByClass(text)
}

// First is a mock method
func (m *AgoutiSelector) First(selector string) *agouti.Selection {
	return m.page.First(selector)
}

// FirstByLink is a mock method
func (m *AgoutiSelector) FirstByLink(text string) *agouti.Selection {
	return m.page.FirstByLink(text)
}

// MockSelectable is a mock of agouti.selectable
type MockSelectable struct {
	mock.Mock
}

// All is a mock method
func (m *MockSelectable) All(selector string) *agouti.MultiSelection {
	args := m.Called(selector)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByXPath is a mock method
func (m *MockSelectable) AllByXPath(selector string) *agouti.MultiSelection {
	args := m.Called(selector)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByLabel is a mock method
func (m *MockSelectable) AllByLabel(text string) *agouti.MultiSelection {
	args := m.Called(text)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByButton is a mock method
func (m *MockSelectable) AllByButton(text string) *agouti.MultiSelection {
	args := m.Called(text)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByName is a mock method
func (m *MockSelectable) AllByName(name string) *agouti.MultiSelection {
	args := m.Called(name)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByLink is a mock method
func (m *MockSelectable) AllByLink(text string) *agouti.MultiSelection {
	args := m.Called(text)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByID is a mock method
func (m *MockSelectable) AllByID(text string) *agouti.MultiSelection {
	args := m.Called(text)
	return args.Get(0).(*agouti.MultiSelection)
}

// AllByClass is a mock method
func (m *MockSelectable) AllByClass(text string) *agouti.MultiSelection {
	args := m.Called(text)
	return args.Get(0).(*agouti.MultiSelection)
}

// First is a mock method
func (m *MockSelectable) First(selector string) *agouti.Selection {
	args := m.Called(selector)
	return args.Get(0).(*agouti.Selection)
}

// FirstByLink is a mock method
func (m *MockSelectable) FirstByLink(text string) *agouti.Selection {
	args := m.Called(text)
	return args.Get(0).(*agouti.Selection)
}

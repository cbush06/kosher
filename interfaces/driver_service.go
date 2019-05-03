package interfaces

import (
	"github.com/sclevine/agouti"
	"github.com/stretchr/testify/mock"
)

// DriverService allows mocking of agouti.WebDriver
type DriverService interface {
	URL() string
	Start() error
	Stop() error
	NewPage(options ...agouti.Option) (*agouti.Page, error)
}

// MockDriver serves as a mock for unit testing
type MockDriver struct {
	mock.Mock
}

// URL is a mock method for unit testing
func (d *MockDriver) URL() string {
	args := d.Called()
	return args.String(0)
}

// Start is a mock method for unit testing
func (d *MockDriver) Start() error {
	args := d.Called()
	return args.Error(0)
}

// Stop is a mock method for unit testing
func (d *MockDriver) Stop() error {
	args := d.Called()
	return args.Error(0)
}

// NewPage is a mock method for unit testing
func (d *MockDriver) NewPage(options ...agouti.Option) (*agouti.Page, error) {
	args := d.Called(options)
	return args.Get(0).(*agouti.Page), args.Error(1)
}

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
	NewPage(options ...agouti.Option) (PageService, error)
}

// UnitTestingMockDriver is the driver returned when unit testing
var UnitTestingMockDriver = new(MockDriver)

// NewDriverService either wraps an agouti.WebDriver (if one is provided) or a MockDriver (if nil is provided) in a DriverService.
func NewDriverService(agoutiDriver *agouti.WebDriver) DriverService {
	if agoutiDriver != nil {
		return &AgoutiDriver{
			driver: agoutiDriver,
		}
	}
	return UnitTestingMockDriver
}

// AgoutiDriver wraps an agouti driver as an adapter
type AgoutiDriver struct {
	driver *agouti.WebDriver
}

// URL returns the URL of the driver endpoint
func (d *AgoutiDriver) URL() string {
	return d.driver.URL()
}

// Start starts the driver
func (d *AgoutiDriver) Start() error {
	return d.driver.Start()
}

// Stop stops the driver
func (d *AgoutiDriver) Stop() error {
	return d.driver.Stop()
}

// NewPage opens a new browser session/page
func (d *AgoutiDriver) NewPage(options ...agouti.Option) (PageService, error) {
	var (
		page *agouti.Page
		err  error
	)

	if page, err = d.driver.NewPage(options...); err != nil {
		return nil, err
	}

	return NewPageServiceFromAgoutiPage(page), nil
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
func (d *MockDriver) NewPage(options ...agouti.Option) (PageService, error) {
	args := d.Called(options)
	return args.Get(0).(PageService), args.Error(1)
}

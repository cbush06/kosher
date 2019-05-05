package clients

import (
	"errors"
	"log"

	"github.com/cbush06/kosher/interfaces"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
)

const (
	// Chrome denotes the Chrome driver
	Chrome = "chrome"

	// Ie denotes the IE driver
	Ie = "ie"

	// Edge denotes the Edge driver
	Edge = "edge"

	// PhantomJS denotes the PhantomJS driver
	PhantomJS = "phantomjs"

	// Desktop denotes the desktop (Appium or WinAppDriver)
	Desktop = "desktop"

	// Mock denotes that a "mock" driver will be used for unit testing
	Mock = "mock"
)

// Client encapsulates the web driver and associated utilities specified by the project's settings.
type Client struct {
	DriverType string
	WebDriver  interfaces.DriverService
	IsStarted  bool
}

// StartDriver starts the driver associated with this client
func (c *Client) StartDriver() error {
	if err := c.WebDriver.Start(); err != nil {
		log.Printf("Kosher failed to connect to driver [%s] at [%s]: %s\n", c.DriverType, c.WebDriver.URL(), err)
		return err
	}
	c.IsStarted = true
	return nil
}

// StopDriver kills the process the driver associated with this client is running on
func (c *Client) StopDriver() error {
	if err := c.WebDriver.Stop(); err != nil {
		log.Printf("Kosher failed to disconnect from driver [%s] at [%s]: %s\n", c.DriverType, c.WebDriver.URL(), err)
		return err
	}
	c.IsStarted = false
	return nil
}

// NewClient is a factory for creating new WebDrivers and associated utilities based on the provided configuration.
func NewClient(sysSettings *config.Settings) (*Client, error) {
	settings := sysSettings.Settings
	driverSetting := settings.GetString("driver")

	switch driverSetting {
	case Chrome:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  interfaces.NewDriverService(agouti.ChromeDriver()),
			IsStarted:  false,
		}, nil
	case Ie:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  interfaces.NewDriverService(agouti.NewWebDriver("http://{{.Address}}", []string{"IEDriverServer.exe", "/port={{.Port}} "})),
			IsStarted:  false,
		}, nil
	case Edge:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  interfaces.NewDriverService(agouti.EdgeDriver()),
			IsStarted:  false,
		}, nil
	case PhantomJS:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  interfaces.NewDriverService(agouti.PhantomJS()),
			IsStarted:  false,
		}, nil
	case Desktop:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  nil,
			IsStarted:  false,
		}, nil
	case Mock:
		return &Client{
			DriverType: driverSetting,
			WebDriver:  interfaces.NewDriverService(nil),
			IsStarted:  false,
		}, nil
	default:
		return nil, errors.New("unknown driver [" + driverSetting + "] specified in settings.json")
	}
}

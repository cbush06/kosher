package clients

import (
	"errors"
	"log"

	"github.com/cbush06/kosher/fs"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
)

// Client encapsulates the web driver and associated utilities specified by the project's settings.
type Client struct {
	DriverType string
	WebDriver  *agouti.WebDriver
}

// StartDriver starts the driver associated with this client
func (c *Client) StartDriver() error {
	if err := c.WebDriver.Start(); err != nil {
		log.Printf("Kosher failed to connect to driver [%s] at [%s]: %s\n", c.DriverType, c.WebDriver.URL(), err)
		return err
	}
	return nil
}

// StopDriver kills the process the driver associated with this client is running on
func (c *Client) StopDriver() error {
	if err := c.WebDriver.Stop(); err != nil {
		log.Printf("Kosher failed to disconnect from driver [%s] at [%s]: %s\n", c.DriverType, c.WebDriver.URL(), err)
		return err
	}
	return nil
}

// NewClient is a factory for creating new WebDrivers and associated utilities based on the provided configuration.
func NewClient(sysSettings *config.Settings, fs *fs.Fs) (*Client, error) {
	settings := sysSettings.Settings
	driverSetting := settings.GetString("driver")

	switch driverSetting {
	case "chrome":
		return &Client{
			DriverType: driverSetting,
			WebDriver:  agouti.ChromeDriver(),
		}, nil
	case "ie":
		return &Client{
			DriverType: driverSetting,
			WebDriver:  agouti.NewWebDriver("http://{{.Address}}", []string{"IEDriverServer.exe", "/port={{.Port}} "}),
		}, nil
	case "edge":
		return &Client{
			DriverType: driverSetting,
			WebDriver:  agouti.EdgeDriver(),
		}, nil
	case "phantomjs":
		return &Client{
			DriverType: driverSetting,
			WebDriver:  agouti.PhantomJS(),
		}, nil
	case "desktop":
		return &Client{
			DriverType: driverSetting,
			WebDriver:  nil,
		}, nil
	default:
		return nil, errors.New("unknown driver [" + driverSetting + "] specified in settings.json")
	}
}

package test

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/cbush06/kosher/clients"
	"github.com/cbush06/kosher/config"
)

func TestNewClient(t *testing.T) {
	var newClientTests = []struct {
		driverType string
		expectNil  bool
		expectErr  bool
	}{
		{"chrome", false, false},
		{"ie", false, false},
		{"edge", false, false},
		{"phantomjs", false, false},
		{"desktop", false, false},
		{"bogus", true, true},
	}

	for _, nextTest := range newClientTests {
		settings := &config.Settings{Settings: viper.New()}
		settings.Settings.Set("driver", nextTest.driverType)

		newClient, err := clients.NewClient(settings)

		if (err != nil && !nextTest.expectErr) || (newClient == nil && !nextTest.expectNil) {
			t.Errorf("failed to create new Client struct for type [%s]: %s", nextTest.driverType, err)
			continue

		} else if (err == nil && nextTest.expectErr) || (newClient != nil && nextTest.expectNil) {
			t.Errorf("expected to fail creating new Client struct for type [%s], but succeeded", nextTest.driverType)
			continue
		}

		if nextTest.expectErr || nextTest.expectNil {
			continue
		}

		if newClient.DriverType != nextTest.driverType {
			t.Errorf("expected driver to be of type [%s], but was of type [%s]", nextTest.driverType, newClient.DriverType)
		}
	}
}

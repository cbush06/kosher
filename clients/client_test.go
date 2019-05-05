package clients

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

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
		t.Run("driverType="+nextTest.driverType, func(t *testing.T) {
			settings := &config.Settings{Settings: viper.New()}
			settings.Settings.Set("driver", nextTest.driverType)

			newClient, err := NewClient(settings)

			if err != nil {
				assert.True(t, nextTest.expectErr, "failed to create new Client struct for type [%s]: %s", nextTest.driverType, err)
				return
			} else if err == nil {
				assert.False(t, nextTest.expectErr, "expected to fail creating new Client struct for type [%s], but succeeded", nextTest.driverType)
				return
			}

			if newClient == nil {
				assert.True(t, nextTest.expectNil, "failed to create new Client struct for type [%s]: %s", nextTest.driverType, err)
				return
			} else if newClient != nil {
				assert.False(t, nextTest.expectNil, "expected to fail creating new Client struct for type [%s], but succeeded", nextTest.driverType)
				return
			}

			if nextTest.expectErr || nextTest.expectNil {
				return
			}

			if newClient.DriverType != nextTest.driverType {
				t.Errorf("expected driver to be of type [%s], but was of type [%s]", nextTest.driverType, newClient.DriverType)
			}
		})
	}
}

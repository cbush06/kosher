package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProvider struct {
	mock.Mock
}

func (m *MockProvider) BindPFlag(key string, flag *pflag.Flag) error {
	args := m.Called(key, flag)
	return args.Error(0)
}

func (m *MockProvider) GetString(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func (m *MockProvider) GetInt(key string) int {
	args := m.Called(key)
	return args.Int(0)
}

func (m *MockProvider) GetBool(key string) bool {
	args := m.Called(key)
	return args.Bool(0)
}

func (m *MockProvider) GetStringMap(key string) map[string]interface{} {
	args := m.Called(key)
	return args.Get(0).(map[string]interface{})
}

func (m *MockProvider) GetStringMapString(key string) map[string]string {
	args := m.Called(key)
	return args.Get(0).(map[string]string)
}

func (m *MockProvider) GetStringSlice(key string) []string {
	args := m.Called(key)
	return args.Get(0).([]string)
}

func (m *MockProvider) Get(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *MockProvider) Set(key string, value interface{}) {
	m.Called(key, value)
}

func (m *MockProvider) IsSet(key string) bool {
	args := m.Called(key)
	return args.Bool(0)
}

func newMockSettings() *Settings {
	return &Settings{
		new(MockProvider),
		new(MockProvider),
		new(MockProvider),
		new(MockProvider),
		nil,
	}
}

func TestGetEnvironmentBaseURL(t *testing.T) {

	t.Run("No-Environment-Setting", func(t *testing.T) {
		settings := newMockSettings()
		settingProv := settings.Settings.(*MockProvider)

		settingProv.On("IsSet", "environment").Return(false)
		actual := settings.GetEnvironmentBaseURL()
		assert.Emptyf(t, actual, "expected empty response, but was [%s]", actual)
	})

	t.Run("Empty-Environment-Setting", func(t *testing.T) {
		settings := newMockSettings()
		settingProv := settings.Settings.(*MockProvider)

		settingProv.On("IsSet", "environment").Return(true)
		settingProv.On("GetString", "environment").Return("")
		actual := settings.GetEnvironmentBaseURL()
		assert.Emptyf(t, actual, "expected empty response, but was [%s]", actual)
	})

	t.Run("Environment-Not-Found", func(t *testing.T) {
		settings := newMockSettings()
		settingProv := settings.Settings.(*MockProvider)
		environmentProv := settings.Environments.(*MockProvider)

		settingProv.On("IsSet", "environment").Return(true)
		settingProv.On("GetString", "environment").Return("dev")
		environmentProv.On("IsSet", "dev").Return(false)
		actual := settings.GetEnvironmentBaseURL()
		assert.Emptyf(t, actual, "expected empty response, but was [%s]", actual)
	})

	t.Run("Environment-Found", func(t *testing.T) {
		settings := newMockSettings()
		settingProv := settings.Settings.(*MockProvider)
		environmentProv := settings.Environments.(*MockProvider)

		settingProv.On("IsSet", "environment").Return(true)
		settingProv.On("GetString", "environment").Return("dev")
		environmentProv.On("IsSet", "dev").Return(true)
		environmentProv.On("GetString", "dev").Return("good to go")
		actual := settings.GetEnvironmentBaseURL()
		assert.Equalf(t, "good to go", actual, "expected [good to go] response, but was [%s]", actual)
	})
}

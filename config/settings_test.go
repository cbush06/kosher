package config

import (
	"bytes"
	"testing"

	"github.com/spf13/afero"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/fs"

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

func TestBuildProvider(t *testing.T) {
	t.Run("File-Not-Found", func(t *testing.T) {
		errorMsg := "Configuration file does not exist"

		fs.MockFs = afero.NewMemMapFs()
		common.BuildTestProject(fs.MockFs)

		fileSys, _ := fs.NewFs("")
		fileSys.ConfigDir.Remove("settings.json")

		var provider Provider
		consoleOutput := common.CaptureStdout(func() {
			provider = buildProvider("settings.json", fileSys, nil)
		})
		assert.Nil(t, provider, "expected provider to  be nil due to non-existent file, but was not")
		assert.Containsf(t, consoleOutput, errorMsg, "expected log message containing [%s], but not found", errorMsg)
	})

	// TODO: Test failure to open file
	// TODO: Test failure to read file

	t.Run("Main-Path", func(t *testing.T) {
		fs.MockFs = afero.NewMemMapFs()
		common.BuildTestProject(fs.MockFs)
		fileSys, _ := fs.NewFs("")

		afero.WriteReader(fileSys.ConfigDir, "settings.json", bytes.NewBufferString(`{ "projectName": "test" }`))

		var provider Provider
		consoleOutput := common.CaptureStdout(func() {
			provider = buildProvider("settings.json", fileSys, nil)
		})
		assert.NotNil(t, provider, "expected provider be successfully created")
		assert.Emptyf(t, consoleOutput, "expected no console output but got [%s]", consoleOutput)
		assert.Equal(t, "test", provider.GetString("projectName"), "expect property [projectName] to equal [test]")
	})

	t.Run("Settings-Provider-Modifier", func(t *testing.T) {
		fs.MockFs = afero.NewMemMapFs()
		common.BuildTestProject(fs.MockFs)
		fileSys, _ := fs.NewFs("")

		var provider Provider
		consoleOutput := common.CaptureStdout(func() {
			provider = buildProvider("settings.json", fileSys, modSettingsProvider)
		})
		assert.NotNil(t, provider, "expected provider be successfully created")
		assert.Emptyf(t, consoleOutput, "expected no console output but got [%s]", consoleOutput)
		assert.Equal(t, "kosher tested app", provider.GetString("projectName"), "expect property [projectName] to equal [kosher tested app]")
	})

	t.Run("Pages-Provider-Modifier", func(t *testing.T) {
		fs.MockFs = afero.NewMemMapFs()
		common.BuildTestProject(fs.MockFs)
		fileSys, _ := fs.NewFs("")

		var provider Provider
		consoleOutput := common.CaptureStdout(func() {
			provider = buildProvider("settings.json", fileSys, modSettingsProvider)
		})
		assert.NotNil(t, provider, "expected provider be successfully created")
		assert.Emptyf(t, consoleOutput, "expected no console output but got [%s]", consoleOutput)
		assert.Equal(t, "kosher tested app", provider.GetString("projectName"), "expect property [projectName] to equal [kosher tested app]")
	})

}

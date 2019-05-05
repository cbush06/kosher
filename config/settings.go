package config

import (
	"bytes"
	"log"
	"os"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/fs"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// Settings wraps references to all setting providers required by Kosher
type Settings struct {
	Environments Provider
	Pages        Provider
	Selectors    Provider
	Settings     Provider
	FileSystem   *fs.Fs
}

// NewSettings attempts to build a Settings object based on the given Fs object
func NewSettings(fs *fs.Fs) *Settings {
	settings := &Settings{
		Environments: buildProvider(common.EnvironmentsFile, fs, nil),
		Pages:        buildProvider(common.PagesFile, fs, modPagesProvider),
		Selectors:    buildProvider(common.SelectorsFile, fs, nil),
		Settings:     buildProvider(common.SettingsFile, fs, modSettingsProvider),
		FileSystem:   fs,
	}
	return settings
}

// GetEnvironmentBaseURL returns the base URL for the environment of the current run
func (s *Settings) GetEnvironmentBaseURL() string {
	if !s.Settings.IsSet("environment") {
		log.Fatal("No setting found for [environment]")
	}
	environment := s.Settings.GetString("environment")
	if !s.Environments.IsSet(environment) {
		log.Fatalf("No entry found for [%s] in the environments file", environment)
	}
	return s.Environments.GetString(environment)
}

type providerModifier func(v *viper.Viper)

func buildProvider(fileName string, fs *fs.Fs, modifyProvider providerModifier) Provider {
	provider := viper.New()
	path, _ := fs.ConfigDir.RealPath(fileName)

	if _, err := fs.ConfigDir.Stat(fileName); err != nil {
		log.Fatalf("Stat failed for [%s]: %s", path, err)
	}
	if exists, err := afero.Exists(fs.ConfigDir, fileName); !exists {
		if err != nil {
			log.Println(err)
		}
		log.Fatal("Configuration file does not exist: " + path)
	}

	if file, err := fs.ConfigDir.OpenFile(fileName, os.O_RDONLY, 0744); err != nil {
		log.Fatalf("Error encountered while opening config file at [%s]: %s", path, err)
	} else {
		if rawBytes, err := afero.ReadAll(file); err != nil {
			log.Fatalf("Error reading config file [%s]: %s", path, err)
		} else {
			provider.SetConfigType("json")
			provider.ReadConfig(bytes.NewBuffer(rawBytes))
		}
	}

	if modifyProvider != nil {
		modifyProvider(provider)
	}

	return provider
}

func modPagesProvider(v *viper.Viper) {
	v.SetDefault("home", "/")
}

func modSettingsProvider(v *viper.Viper) {
	v.SetDefault("cucumberDialect", "en")
	v.SetDefault("projectName", "kosher tested app")
	v.SetDefault("appVersion", "1.0.0")
	v.SetDefault("platform", "web")
	v.SetDefault("driver", "chrome")
	v.SetDefault("reportFormat", "pretty")
	v.SetDefault("dateFormat", "MM/DD/YYYY")
	v.SetDefault("defaultEnvironment", "test")
	v.SetDefault("screenFormat", "desktop")
	v.SetDefault("quitOnFail", false)
	v.SetDefault("ignoreInvisible", true)
	v.SetDefault("screenFormats.desktop.width", 2000)
	v.SetDefault("screenFormats.desktop.height", 980)
	v.SetDefault("screenFormats.mobile.width", 362)
	v.SetDefault("screenFormats.mobile.height", 868)
	v.SetDefault("screenFormats.tablet.width", 814)
	v.SetDefault("screenFormats.tablet.height", 868)
	v.SetDefault("screenFormats.landscape.width", 522)
	v.SetDefault("screenFormats.landscape.height", 362)
	v.SetDefault("integrations.jira.host", "http://127.0.0.1")
	v.SetDefault("integrations.jira.defaults.projectKey", "PROJE")
	v.SetDefault("integrations.jira.defaults.issueType", "Bug")
	v.SetDefault("integrations.jira.defaults.affectsVersion", "1.0.0")
	v.SetDefault("integrations.jira.defaults.labels", "test,functional,kosher")
	v.SetDefault("integrations.jira.defaults.priority", "Normal")
}

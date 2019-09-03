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
		log.Print("No setting found for [environment]\n")
		return ""
	}

	environment := s.Settings.GetString("environment")
	if len(environment) < 1 {
		log.Print("Empty setting found for [environment]\n")
		return ""
	}

	if !s.Environments.IsSet(environment) {
		log.Printf("No entry found for [%s] in the environments file\n", environment)
		return ""
	}
	return s.Environments.GetString(environment)
}

type providerModifier func(v *viper.Viper)

func buildProvider(fileName string, fs *fs.Fs, modifyProvider providerModifier) Provider {
	var (
		provider = viper.New()
		path, _  = fs.ConfigDir.RealPath(fileName)
		exists   bool
		rawBytes []byte
		file     afero.File
		err      error
	)

	if exists, err = afero.Exists(fs.ConfigDir, fileName); !exists || err != nil {
		if err != nil {
			log.Println(err)
		}
		log.Println("Configuration file does not exist: " + path)
		return nil
	}

	if file, err = fs.ConfigDir.OpenFile(fileName, os.O_RDONLY, 0744); err != nil {
		log.Printf("Error encountered while opening config file at [%s]: %s\n", path, err)
		return nil
	}

	if rawBytes, err = afero.ReadAll(file); err != nil {
		log.Printf("Error reading config file [%s]: %s\n", path, err)
		return nil
	}

	provider.SetConfigType("json")
	provider.ReadConfig(bytes.NewBuffer(rawBytes))

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
	v.SetDefault("waitAfterScenario", 0)
	v.SetDefault("waitAfterStep", 0)
	v.SetDefault("accessibility.ruleSets", []string{"wcag21aa", "section508"})
	v.SetDefault("accessibility.impactThreshold", "critical")
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

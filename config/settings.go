package config

import (
	"fmt"
	"log"

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
}

// NewSettings attempts to build a Settings object based on the given Fs object
func NewSettings(fs *fs.Fs) *Settings {
	settings := &Settings{
		Environments: buildProvider(common.EnvironmentsFile, fs, nil),
		Pages:        buildProvider(common.PagesFile, fs, modPagesProvider),
		Selectors:    buildProvider(common.SelectorsFile, fs, nil),
		Settings:     buildProvider(common.SettingsFile, fs, modSettingsProvider),
	}
	return settings
}

type providerModifier func(v *viper.Viper)

func buildProvider(fileName string, fs *fs.Fs, modifyProvider providerModifier) Provider {
	provider := viper.New()
	path, _ := fs.ConfigDir.RealPath(fileName)

	if _, err := fs.ConfigDir.Stat(fileName); err != nil {
		log.Fatal("Stat failed for [" + fileName + "]: " + err.Error())
	}
	if exists, err := afero.Exists(fs.ConfigDir, fileName); !exists {
		if err != nil {
			fmt.Println(err.Error())
		}
		log.Fatal("Configuration file does not exist: " + path)
	}

	provider.SetConfigFile(path)
	if modifyProvider != nil {
		modifyProvider(provider)
	}
	return provider
}

func modPagesProvider(v *viper.Viper) {
	v.SetDefault("home", "/")
}

func modSettingsProvider(v *viper.Viper) {
	v.SetDefault("defaultDateFormat", "mm/dd/yyyy")
	v.SetDefault("defaultEnvironment", "test")
	v.SetDefault("maxWaitTime", 2)
	v.SetDefault("minWaitTime", 0)
	v.SetDefault("httpTimeout", 3)
	v.SetDefault("waitForPageLoad", 0)
	v.SetDefault("screenSize", "desktop")
	v.SetDefault("debugMode", "")
	v.SetDefault("quitOnFail", false)
	v.SetDefault("screenSizeConfigurations", map[string]map[string]int{
		"desktop": map[string]int{
			"width":  2000,
			"height": 980,
		},
		"mobile": map[string]int{
			"width":  362,
			"height": 868,
		},
		"tablet": {
			"width":  814,
			"height": 868,
		},
		"landscape": {
			"width":  522,
			"height": 362,
		},
	})
}

package config

import (
	"github.com/spf13/pflag"
)

// Provider provides the configuration settings for Kosher
type Provider interface {
	BindPFlag(key string, flag *pflag.Flag) error
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	Get(key string) interface{}
	Set(key string, value interface{})
	IsSet(key string) bool
}

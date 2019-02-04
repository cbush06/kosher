package config

// Provider provides the configuration settings for Kosher
type Provider interface {
	GetString(key string) string
}

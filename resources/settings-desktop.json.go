package resources

// GetSettingsDesktopJSON returns the settings.json template file
func GetSettingsDesktopJSON() string {
	return `{
	"driver": "desktop",
	"appId": "<appId for desktop application>",
	"deviceName": "PC name"
}`
}

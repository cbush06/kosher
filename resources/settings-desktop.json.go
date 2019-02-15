package resources

// GetSettingsDesktopJSON returns the settings.json template file
func GetSettingsDesktopJSON() string {
	return `{
	"projectName": "kosher",
	"platform": "desktop",
	"driver": "appium",
	"reportFormat": "pretty",
	"dateFormat": "MM/DD/YYYY",
	"appId": "<appId for desktop application>",
	"deviceName": "PC name"
}`
}
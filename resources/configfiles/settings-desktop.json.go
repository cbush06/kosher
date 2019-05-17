package configfiles

// GetSettingsDesktopJSON returns the settings.json template file
func GetSettingsDesktopJSON(projectName string) string {
	return `{
	"cucumberDialect": "en",
    "projectName": "` + projectName + `",
    "appVersion": "1.0.0",
    "platform": "desktop",
    "driver": "appium",
    "reportFormat": "pretty",
    "dateFormat": "MM/DD/YYYY",
    "appId": "<appId for desktop application>",
	"deviceName": "PC name",
	"quitOnFail": false,
	"ignoreInvisible": true,
	"waitAfterScenario": 0,
	"waitAfterStep": 0,
	"integrations": {
		"jira": {
			"host": "http://127.0.0.1",
			"defaults": {
				"projectKey": "PROJE",
				"issueType": "Bug",
				"affectsVersion": "1.0.0",
				"labels": "test,functional,kosher",
				"priority": "Normal"
			}
		}
	}
}`
}

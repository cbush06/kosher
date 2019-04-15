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
	"integrations": {
		"jira": {
			"host": "https://127.0.0.1",
			"labels": "test,functional,kosher",
			"summaryTemplate": "jira_summary.txt",
			"descriptionTemplate": "jira_description.txt",
			"acceptanceCriteriaTemplate": "jira_acceptancecriteria.txt"
		}
	}
}`
}

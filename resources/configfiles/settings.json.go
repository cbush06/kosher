package configfiles

// GetSettingsJSON returns the settings.json template file
func GetSettingsJSON(projectName string) string {
	return `{
	"cucumberDialect": "en",
    "projectName": "` + projectName + `",
    "appVersion": "1.0.0",
    "platform": "web",
    "driver": "chrome",
    "reportFormat": "pretty",
    "dateFormat": "MM/DD/YYYY",
    "defaultEnvironment": "test",
    "screenFormat": "desktop",
	"quitOnFail": false,
	"ignoreInvisible": true,
    "screenFormats": {
        "desktop": {
            "width": 2000,
            "height": 980
         },
        "mobile": {
            "width": 362,
            "height": 868
        },
        "tablet": {
            "width": 814,
            "height": 868
        },
        "landscape": {
            "width": 522,
            "height": 362
        }
	},
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

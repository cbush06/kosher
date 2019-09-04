package configfiles

// GetSettingsJSON returns the settings.json template file
func GetSettingsJSON(projectName string) string {
	return `{
	"cucumberDialect": "en",
    "projectName": "` + projectName + `",
    "appVersion": "1.0.0",
    "driver": "chrome",
    "reportFormat": "pretty",
    "dateFormat": "MM/DD/YYYY",
    "defaultEnvironment": "test",
    "screenFormat": "desktop",
	"quitOnFail": false,
	"ignoreInvisible": true,
	"waitAfterScenario": 0,
	"waitAfterStep": 0,
	"accessibility": {
		"ruleSets": [
			"wcag21aa",
			"section508"
		],
		"impactThreshold": "serious"
	},
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

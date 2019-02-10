package resources

// GetSettingsJSON returns the settings.json template file
func GetSettingsJSON() string {
	return `{
	"projectName": "kosher",
	"platform": "web",
	"driver": "chrome",
	"reportFormat": "pretty",
    "defaultEnvironment": "test",
    "maxWaitTime": 2,
    "minWaitTime": 0,
    "httpTimeout": 3,
    "waitForPageLoad": 0,
    "screenFormat": "desktop",
    "debugMode": "",
    "quitOnFail": false,
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
    }
}`
}

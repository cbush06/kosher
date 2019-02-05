package resources

// GetEnvironmentsJSON returns the environments.json template file
func GetEnvironmentsJSON() string {
	return `
{
	"production": "http://www.your-production-env.com/",
	"test": "http://www.your-test-env.com",
	"dev": "http://www.your-dev-env.com",
	"uat": "http://www.your-uat-env.com"
}
	`
}

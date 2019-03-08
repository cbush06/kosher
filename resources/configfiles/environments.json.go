package configfiles

// GetEnvironmentsJSON returns the environments.json template file
func GetEnvironmentsJSON() string {
	return `{
	"production": "http://www.your-production-env.com/",
	"test": "https://www.seleniumeasy.com/test",
	"dev": "http://www.your-dev-env.com",
	"uat": "http://www.your-uat-env.com"
}`
}

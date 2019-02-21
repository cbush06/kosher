package resources

// GetPagesJSON returns the pages.json template file
func GetPagesJSON() string {
	return `{
	"home": "",
	"registration": "user/new_user",
	"profile": "user/profile",
	"contact": "forms/contact"
}`
}

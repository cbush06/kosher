package configfiles

// GetPagesJSON returns the pages.json template file
func GetPagesJSON() string {
	return `{
    "home": "/",
    "basicform": "basic-first-form-demo.html",
    "table-sort": "table-sort-search-demo.html"
}`
}

package resources

// GetSelectorsJSON returns the selectors.json template file
func GetSelectorsJSON() string {
	return `
{
	"header1Css": "css:html > body > table > tbody > tr > th:first-child",
	"header2Xpath": "xpath:/html/body/table/tbody/tr[1]/th[2]"
}
	`
}

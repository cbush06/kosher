package steputils

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/cbush06/kosher/clients"
	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/interfaces"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

var (
	flagRegEx = regexp.MustCompile(`^\@\{(.+)\}(.+)$`)
)

// StepUtils is a set of utility functions tailored to a given Settings object and Agouti Page
type StepUtils struct {
	Settings *config.Settings
	Page     interfaces.PageService
	Session  interfaces.SessionService
}

// NewStepUtils creates a new StepUtils struct
func NewStepUtils(settings *config.Settings, page interfaces.PageService) *StepUtils {
	var session interfaces.SessionService = page.Session()

	// If we're using a nil driver for unit testing, sub in the MockSession
	if settings.Settings.GetString("driver") == clients.Mock && session == nil {
		session = &interfaces.MockSession{}
	}

	return &StepUtils{
		Settings: settings,
		Page:     page,
		Session:  session,
	}
}

// ResolvePage takes a page name and returns its full URL
func (s *StepUtils) ResolvePage(pageName string) (string, error) {
	if !s.Settings.Pages.IsSet(pageName) {
		return "", fmt.Errorf("no entry found in the pages file for [%s]", pageName)
	}

	// get the environment base url
	baseURL, _ := url.Parse(s.Settings.GetEnvironmentBaseURL())

	// get the URL of the specified page
	pageURL := s.Settings.Pages.GetString(pageName)

	// join the base URL and page path together
	baseURL.Path = path.Join(baseURL.Path, pageURL)

	return baseURL.String(), nil
}

// ResolveSelector attempts to retrieve the selector specified by `name` and convert it into an Agouti selector for the provided page.
// If no selector is found, `name` is used to search by label, name, and ID (in that order).
func (s *StepUtils) ResolveSelector(name string) ([]*agouti.Selection, error) {
	var (
		ignoreInvisible = s.Settings.Settings.GetBool("ignoreInvisible")
		results         []*agouti.Selection
	)

	// parse flags
	matches := flagRegEx.FindStringSubmatch(name)
	if len(matches) > 0 {
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid selector flag(s)")
		}

		// split flags and process each
		flags := strings.Split(matches[1], ",")
		for _, flag := range flags {
			switch strings.ToUpper(flag) {
			case "INVISIBLE":
				ignoreInvisible = false
			default:
				return nil, fmt.Errorf("unrecognized flag [%s]", flag)
			}
		}

		name = matches[2]
	}

	selector := s.Settings.Selectors.GetString(name)
	if len(selector) > 0 {
		// the name matched an entry in the selectors file, so use that to query
		firstColonIdx := strings.Index(selector, ":")
		selectorType := selector[:firstColonIdx]
		selectorBody := selector[firstColonIdx+1:]

		switch selectorType {
		case "css":
			agoutiSel := s.Page.All(selectorBody)
			matchCnt, _ := agoutiSel.Count()
			if matchCnt > 0 {
				for i := 0; i < matchCnt; i++ {
					if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
						results = append(results, agoutiSel.At(i))
					}
				}
			}
			break
		case "xpath":
			agoutiSel := s.Page.AllByXPath(selectorBody)
			matchCnt, _ := agoutiSel.Count()
			if matchCnt > 0 {
				for i := 0; i < matchCnt; i++ {
					if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
						results = append(results, agoutiSel.At(i))
					}
				}
			}
			break
		default:
			return nil, fmt.Errorf(`invalid selector type [%s] specified for selector [%s] (expected "css:" or "xpath:")`, selectorType, name)
		}

		if len(results) > 0 {
			return results, nil
		}
	} else {
		selector = strings.TrimSpace(name)

		// try to find a match by label
		agoutiSel := s.Page.AllByLabel(selector)
		matchCnt, _ := agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		// try to find a button
		agoutiSel = s.Page.AllByButton(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		// try to find a match by link text
		agoutiSel = s.Page.AllByLink(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		// try to find a match by name
		agoutiSel = s.Page.AllByName(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		// try to find a match by id
		agoutiSel = s.Page.AllByID(name)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		// try to find by text
		agoutiSel = s.Page.AllByXPath(fmt.Sprintf(`//*[contains(text(), '%s')]`, selector))
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			for i := 0; i < matchCnt; i++ {
				if visible, _ := agoutiSel.At(i).Visible(); visible || !ignoreInvisible {
					results = append(results, agoutiSel.At(i))
				}
			}
		}

		if len(results) > 0 {
			return results, nil
		}
	}
	return nil, fmt.Errorf("no matches found for field [%s]", name)
}

// GetMaxWindowSize attempts to determine and return the window size set
func (s *StepUtils) GetMaxWindowSize() (int, int) {
	var (
		screenFormat string
		width        int
		height       int
	)

	if screenFormat = s.Settings.Settings.GetString("screenFormat"); len(screenFormat) < 1 {
		screenFormat = common.DefaultScreenFormat
	}

	if !s.Settings.Settings.IsSet("screenFormats." + screenFormat) {
		width = common.DefaultScreenWidth
		height = common.DefaultScreenHeight
	} else {
		width = s.Settings.Settings.GetInt("screenFormats." + screenFormat + ".width")
		height = s.Settings.Settings.GetInt("screenFormats." + screenFormat + ".height")
	}

	if width < 1 || height < 1 {
		log.Fatalf("invalid screen width and/or height specified")
	}

	return width, height
}

// GetFieldType attempts to determine which type of form field an element is. `name` is just the string
// used to identify the field in the user's gherkin script. This may be the label, name, or ID of the field;
// it, also, may be a key to a selector entry in the selectors file.
func (s *StepUtils) GetFieldType(name string, sel *agouti.Selection) (string, error) {
	var (
		elms      []*api.Element
		tagName   string
		inputType string
		err       error
		errorMsg  = fmt.Sprintf("error encountered while trying to determine tag type of [%s]", name)
	)

	// try to get the actual webdriver element
	if elms, err = sel.Elements(); err != nil {
		return "", fmt.Errorf(errorMsg+": %s", err)
	} else if len(elms) < 1 {
		return "", fmt.Errorf(errorMsg+": no matches found for field [%s]", name)
	}

	// try to get the elements tag name
	if tagName, err = elms[0].GetName(); err != nil {
		return "", fmt.Errorf(errorMsg+"--it's likely the element was found: %s", err)
	}

	// if it's an input field, return its input type; otherwise, return the tag name
	switch strings.ToLower(tagName) {
	case "input":
		if inputType, err = elms[0].GetAttribute("type"); err != nil {
			return "", fmt.Errorf(errorMsg+": %s", err)
		}
		return inputType, nil
	default:
		return tagName, nil
	}
}

// FormatDate converts a `time.Time` to a string using the `dateFormat` value specified in the settings file
func (s *StepUtils) FormatDate(t time.Time) string {
	dateFormat := s.Settings.Settings.GetString("dateFormat")
	return common.FormatDate(t, dateFormat)
}

// ParseDate parses a given string to a `time.Time` using the `dateFormat` value specified in the settings file
func (s *StepUtils) ParseDate(date string) time.Time {
	dateFormat := s.Settings.Settings.GetString("dateFormat")
	return common.ParseDate(date, dateFormat)
}

// IsFormField determines if the selection provided is an HTML form field
func (s *StepUtils) IsFormField(field string, sel *agouti.Selection) bool {
	var (
		fieldType string
		err       error
	)

	if fieldType, err = s.GetFieldType(field, sel); err != nil {
		return false
	}

	switch fieldType {
	case "date", "datetime-local", "email", "month", "number", "password", "search", "tel", "time", "url", "week", "textarea", "text", "button", "submit",
		"image", "reset", "radio", "checkbox", "select", "range", "file", "color":
		return true
	default:
		return false
	}
}

// IsTextBased determines if a given field is a form of textbox or textarea
func (s *StepUtils) IsTextBased(field string, sel *agouti.Selection) bool {
	var (
		fieldType string
		err       error
	)

	if fieldType, err = s.GetFieldType(field, sel); err != nil {
		return false
	}

	switch fieldType {
	case "date", "datetime-local", "email", "month", "number", "password", "search", "tel", "time", "url", "week", "textarea", "text":
		return true
	default:
		return false
	}
}

// GetSelectOptions returns a map where the keys are the visible text of each option and
// the values are bools indicating the selected statuses of those options
func (s *StepUtils) GetSelectOptions(htmlSelect *agouti.Selection) map[string]bool {
	// get the options
	options := htmlSelect.All("option")
	optionsElms, _ := options.Elements()

	// map the options to (text, bool) pairs
	var results = make(map[string]bool)
	for _, nextElm := range optionsElms {
		text, _ := nextElm.GetText()
		selected, _ := nextElm.IsSelected()
		results[strings.TrimSpace(text)] = selected
	}

	return results
}

// ReplaceVariables interpolates variables found in step arguments
func (s *StepUtils) ReplaceVariables(text string) (string, error) {
	regEx := regexp.MustCompile(`\$\{(.+)\}`)
	matches := regEx.FindAllStringSubmatch(text, -1)
	interpolatedText := text

	for _, nextMatch := range matches {
		switch variable := strings.ToUpper(nextMatch[1]); variable {
		case "BACKSPACE", "ENTER", "ESCAPE", "SPACE", "DELETE":
			interpolatedText = strings.Replace(interpolatedText, nextMatch[0], s.getKeyCode(variable), -1)
		case "RESOURCESDIR":
			interpolatedText = strings.Replace(interpolatedText, nextMatch[0], s.getPath(variable), -1)
		default:
			return "", fmt.Errorf("unrecognized variable [%s]", nextMatch[1])
		}
	}

	return interpolatedText, nil
}

func (s *StepUtils) getPath(variable string) string {
	var path string

	switch variable {
	case "RESOURCESDIR":
		path, _ = s.Settings.FileSystem.WorkingDir.RealPath(common.ResourcesDir)
	default:
		return ""
	}
	return path
}

// Borrowed these unicodes from https://github.com/SeleniumHQ/selenium/blob/master/java/client/src/org/openqa/selenium/Keys.java
func (s *StepUtils) getKeyCode(variable string) string {
	switch variable {
	case "BACKSPACE":
		return string(0xE003)
	case "ENTER":
		return string(0xE007)
	case "ESCAPE":
		return string(0xE00C)
	case "SPACE":
		return string(0xE00D)
	case "DELETE":
		return string(0xE017)
	default:
		return ""
	}
}

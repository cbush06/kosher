package steputils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cbush06/kosher/common"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

// StepUtils is a set of utility functions tailored to a given Settings object and Agouti Page
type StepUtils struct {
	Settings *config.Settings
	Page     *agouti.Page
}

// NewStepUtils creates a new StepUtils struct
func NewStepUtils(settings *config.Settings, page *agouti.Page) *StepUtils {
	return &StepUtils{
		Settings: settings,
		Page:     page,
	}
}

// ResolveSelector attempts to retrieve the selector specified by `name` and convert it into an Agouti selector for the provided page.
// If no selector is found, `name` is used to search by label, name, and ID (in that order).
func (s *StepUtils) ResolveSelector(name string) (*agouti.MultiSelection, error) {
	var selector = s.Settings.Selectors.GetString(name)

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
				return agoutiSel, nil
			}
			break
		case "xpath":
			agoutiSel := s.Page.AllByXPath(selectorBody)
			matchCnt, _ := agoutiSel.Count()
			if matchCnt > 0 {
				return agoutiSel, nil
			}
			break
		default:
			return nil, fmt.Errorf("invalid selector type [%s] specified for selector [%s]", selectorType, name)
		}
	} else {
		selector = strings.TrimSpace(name)

		// try to find a match by label
		agoutiSel := s.Page.AllByLabel(selector)
		matchCnt, _ := agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a button
		agoutiSel = s.Page.AllByButton(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a match by link text
		agoutiSel = s.Page.AllByLink(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a match by name
		agoutiSel = s.Page.AllByName(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a match by id
		agoutiSel = s.Page.AllByID(name)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find by text
		agoutiSel = s.Page.AllByXPath(fmt.Sprintf(`//*[contains(text(), '%s')]`, selector))
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
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
	if len(dateFormat) < 1 {
		return ""
	}
	return t.Format(convertDateFormatToGoFormat(dateFormat))
}

// ParseDate parses a given string to a `time.Time` using the `dateFormat` value specified in the settings file
func (s *StepUtils) ParseDate(date string) time.Time {
	dateFormat := s.Settings.Settings.GetString("dateFormat")
	if len(dateFormat) < 1 {
		return time.Time{}
	}
	if goDate, err := time.ParseInLocation(convertDateFormatToGoFormat(dateFormat), date, time.Local); err != nil {
		return time.Time{}
	} else {
		return goDate
	}
}

func convertDateFormatToGoFormat(dateFormat string) string {
	dateFormat = strings.Replace(dateFormat, "MMMM", "January", 1)
	dateFormat = strings.Replace(dateFormat, "MMM", "Jan", 1)
	dateFormat = strings.Replace(dateFormat, "MM", "01", 1)
	dateFormat = strings.Replace(dateFormat, "YYYY", "2006", 1)
	dateFormat = strings.Replace(dateFormat, "YY", "06", 1)
	dateFormat = strings.Replace(dateFormat, "DDDD", "Monday", 1)
	dateFormat = strings.Replace(dateFormat, "DDD", "Mon", 1)
	dateFormat = strings.Replace(dateFormat, "DD", "02", 1)
	return dateFormat
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

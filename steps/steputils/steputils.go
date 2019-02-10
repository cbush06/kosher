package steputils

import (
	"fmt"
	"log"
	"strings"

	"github.com/cbush06/kosher/common"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

// StepUtils is a set of utility functions tailored to a given Settings object and Agouti Page
type StepUtils struct {
	settings *config.Settings
	page     *agouti.Page
}

// NewStepUtils creates a new StepUtils struct
func NewStepUtils(settings *config.Settings, page *agouti.Page) *StepUtils {
	return &StepUtils{
		settings: settings,
		page:     page,
	}
}

// ResolveSelector attempts to retrieve the selector specified and convert it into an Agouti selector for the provided page.
// If no selector is found, the value is used to search by label, name, and ID (in that order).
func (s *StepUtils) ResolveSelector(name string) (*agouti.MultiSelection, error) {
	var selector = s.settings.Selectors.GetString(name)

	if len(selector) > 0 {
		// the name matched an entry in the selectors file, so use that to query
		firstColonIdx := strings.Index(selector, ":")
		selectorType := selector[:firstColonIdx]
		selectorBody := selector[firstColonIdx+1:]

		switch selectorType {
		case "css":
			agoutiSel := s.page.All(selectorBody)
			matchCnt, _ := agoutiSel.Count()
			if matchCnt > 0 {
				return agoutiSel, nil
			}
			break
		case "xpath":
			agoutiSel := s.page.AllByXPath(selectorBody)
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
		agoutiSel := s.page.AllByLabel(selector)
		matchCnt, _ := agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a match by name
		agoutiSel = s.page.AllByName(selector)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}

		// try to find a match by id
		agoutiSel = s.page.AllByID(name)
		matchCnt, _ = agoutiSel.Count()
		if matchCnt > 0 {
			return agoutiSel, nil
		}
	}
	return nil, fmt.Errorf("no matches found for [%s]", name)
}

// GetMaxWindowSize attempts to determine and return the window size set
func (s *StepUtils) GetMaxWindowSize() (int, int) {
	var (
		screenFormat string
		width        int
		height       int
	)

	if screenFormat = s.settings.Settings.GetString("screenFormat"); len(screenFormat) < 1 {
		screenFormat = common.DefaultScreenFormat
	}

	if !s.settings.Settings.IsSet("screenFormats." + screenFormat) {
		width = common.DefaultScreenWidth
		height = common.DefaultScreenHeight
	} else {
		width = s.settings.Settings.GetInt("screenFormats." + screenFormat + ".width")
		height = s.settings.Settings.GetInt("screenFormats." + screenFormat + ".height")
	}

	if width < 1 || height < 1 {
		log.Fatalf("invalid screen width and/or height specified")
	}

	return width, height
}

// GetFieldType attempts to determine which type of form field an element is. selOrLabel is just the string
// used to identify the field in the user's gherkin script. This may be the label of the field or a key
// to a selector entry in the selectors file.
func (s *StepUtils) GetFieldType(selOrLabel string, sel *agouti.Selection) (string, error) {
	var (
		elms      []*api.Element
		tagName   string
		inputType string
		err       error
		errorMsg  = fmt.Sprintf("error encountered while trying to determine tag type of [%s]", selOrLabel)
	)

	// try to get the actual webdriver element
	if elms, err = sel.Elements(); err != nil {
		return "", fmt.Errorf(errorMsg+": %s", err)
	} else if len(elms) < 1 {
		return "", fmt.Errorf(errorMsg + ": no matches found")
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

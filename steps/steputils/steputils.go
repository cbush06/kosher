package steputils

import (
	"fmt"
	"log"
	"strings"

	"github.com/cbush06/kosher/common"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
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

// ResolveSelector attempts to retrieve the selector specified and convert it into an Agouti selector for the provided page
func (s *StepUtils) ResolveSelector(name string) (*agouti.Selection, error) {
	var selector string
	if selector = s.settings.Selectors.GetString(name); len(selector) < 1 {
		return nil, fmt.Errorf("no selector found for name [%s]", name)
	}

	firstColonIdx := strings.Index(selector, ":")
	selectorType := selector[:firstColonIdx]
	selectorBody := selector[firstColonIdx+1:]

	switch selectorType {
	case "css":
		return s.page.Find(selectorBody), nil
	case "xpath":
		return s.page.FindByXPath(selectorBody), nil
	default:
		return nil, fmt.Errorf("invalid selector type [%s] specified for selector [%s]", selectorType, name)
	}
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

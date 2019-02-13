package websteps

import (
	"fmt"

	"github.com/cbush06/kosher/steps/steputils"
)

func iShouldBeRedirectedTo(s *steputils.StepUtils) func(string) error {
	return iShouldBeOn(s)
}

func iShouldBeOn(s *steputils.StepUtils) func(string) error {
	return func(pageName string) error {
		var (
			expectedURL string
			err         error
		)

		if expectedURL, err = s.ResolvePage(pageName); err != nil {
			return fmt.Errorf("error encountered while verifying page URL: %s", err)
		}

		// get this page's URL
		currentURL, _ := s.Page.URL()

		// assert their equality
		if expectedURL != currentURL {
			return fmt.Errorf("expected URL to be [%s] but was [%s]", expectedURL, currentURL)
		}
		return nil
	}
}

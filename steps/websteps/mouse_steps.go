package websteps

import (
	"fmt"

	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func iHoverOver(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		var (
			matches []*agouti.Selection
			errMsg  = fmt.Sprintf("error encountered while hovering over [%s]: ", field) + "%s"
			err     error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's at least 1
		fieldCnt := len(matches)
		if fieldCnt < 0 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// hover over it
		if err = matches[0].MouseToElement(); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		return nil
	}
}

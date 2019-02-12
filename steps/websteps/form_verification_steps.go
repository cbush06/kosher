package websteps

import (
	"fmt"
	"log"
	"time"

	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func iVerifyHasTodaysDate(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		var (
			matches *agouti.MultiSelection
			errMsg  = fmt.Sprintf("error encountered while verifying today's date is in [%s]: ", field) + "%s"
			err     error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's at least 1
		fieldCnt, _ := matches.Count()
		if fieldCnt < 0 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// get the value
		fieldElms, _ := matches.At(0).Elements()
		fieldVal, _ := fieldElms[0].GetAttribute("value")

		// convert it to a `time.Time`
		fieldValTime := s.ParseDate(fieldVal)

		// verify it's equal to today
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

		if fieldValTime.Sub(today) != 0 {
			log.Println(fieldValTime.Sub(today))
			return fmt.Errorf("expected to find today's date [%s] in [%s] but found [%s] instead", s.FormatDate(now), field, fieldVal)
		}

		return nil
	}
}

func shouldContain(s *steputils.StepUtils) func(string, string) error {
	return shouldContainNotContain(s, true)
}

func shouldNotContain(s *steputils.StepUtils) func(string, string) error {
	return shouldContainNotContain(s, false)
}

func shouldContainNotContain(s *steputils.StepUtils, shouldContain bool) func(string, string) error {
	return func(field string, value string) error {
		var (
			matches *agouti.MultiSelection
			errMsg  = fmt.Sprintf("error encountered while verifying today's date is in [%s]: ", field) + "%s"
			err     error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's at least 1
		fieldCnt, _ := matches.Count()
		if fieldCnt < 0 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// get the value
		if s.IsTextBased(field, matches.At(0)) {
			fieldElms, _ := matches.At(0).Elements()
			fieldVal, _ := fieldElms[0].GetAttribute("value")
			doesMatch := value == fieldVal

			if doesMatch {
				if shouldContain {
					return nil
				}
				return fmt.Errorf("expected [%s] to contain [%s] but it contained [%s]", field, value, fieldVal)
			}

			if !shouldContain {
				return nil
			}
			return fmt.Errorf("expected [%s] to NOT contain [%s] but it contained [%s]", field, value, fieldVal)
		}
		return fmt.Errorf(errMsg, "field must be some form of textbox")
	}
}

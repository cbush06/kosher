package websteps

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func shouldContainTodaysDate(s *steputils.StepUtils) func(string) error {
	return iVerifyHasTodaysDate(s)
}

func iVerifyHasTodaysDate(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		var (
			matches []*agouti.Selection
			errMsg  = fmt.Sprintf("error encountered while verifying today's date is in [%s]: ", field) + "%s"
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

		// get the value
		fieldElms, _ := matches[0].Elements()
		fieldVal, _ := fieldElms[0].GetAttribute("value")

		// convert it to a `time.Time`
		fieldValTime := s.ParseDate(fieldVal)

		// verify it's equal to today
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

		if fieldValTime.Sub(today) != 0 {
			return fmt.Errorf("expected to find today's date [%s] in [%s] but found [%s] instead", s.FormatDate(now), field, fieldVal)
		}

		return nil
	}
}

func shouldContain(s *steputils.StepUtils) func(string, string) error {
	return confirmContents(s, true)
}

func shouldNotContain(s *steputils.StepUtils) func(string, string) error {
	return confirmContents(s, false)
}

func confirmContents(s *steputils.StepUtils, shouldContain bool) func(string, string) error {
	return func(field string, value string) error {
		var (
			matches []*agouti.Selection
			errMsg  = fmt.Sprintf("error encountered while verifying contents of [%s]: ", field) + "%s"
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

		// get the value
		if s.IsTextBased(field, matches[0]) {
			fieldElms, _ := matches[0].Elements()
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

func shouldHaveTheFollowingOptions(s *steputils.StepUtils) func(string, *gherkin.DataTable) error {
	return confirmSelectOptions(s, false, false)
}

func shouldHaveTheFollowingOptionsSelected(s *steputils.StepUtils) func(string, *gherkin.DataTable) error {
	return confirmSelectOptions(s, true, true)
}

func shouldNotHaveTheFollowingOptionsSelected(s *steputils.StepUtils) func(string, *gherkin.DataTable) error {
	return confirmSelectOptions(s, true, false)
}

func confirmSelectOptions(s *steputils.StepUtils, only bool, selected bool) func(string, *gherkin.DataTable) error {
	return func(field string, expectedOptions *gherkin.DataTable) error {
		var (
			matches    []*agouti.Selection
			errMsg     = fmt.Sprintf("error encountered while verifying options for [%s]: ", field) + "%s"
			noMatchMsg = "actual options differed from expected options"
			err        error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's at least 1
		fieldCnt := len(matches)
		if fieldCnt < 1 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// ensure its a select
		if fieldType, err := s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, fmt.Sprintf("error encountered determining field type: %s", err))
		} else if !strings.EqualFold(fieldType, "select") {
			return fmt.Errorf(errMsg, "[%s] must be of type [select] but is [%s]", field, fieldType)
		}

		// get option elements
		optionElms := s.GetSelectOptions(matches[0])

		// map expected values to a slice
		var expected []string
		for _, nextRow := range expectedOptions.Rows {
			expected = append(expected, strings.TrimSpace(nextRow.Cells[0].Value))
		}
		sort.Strings(expected)

		// map actual values to a slice
		var actual []string
		for optionText, isSelected := range optionElms {
			if (only && ((selected && isSelected) || !(selected || isSelected))) || !only {
				actual = append(actual, optionText)
			}
		}
		sort.Strings(actual)

		// ensure they're equal
		if len(expected) != len(actual) {
			return fmt.Errorf(errMsg, noMatchMsg)
		}
		for i := 0; i < len(expected); i++ {
			if actual[i] != expected[i] {
				return fmt.Errorf(errMsg, noMatchMsg)
			}
		}

		return nil
	}
}

func shouldBeSelected(s *steputils.StepUtils) func(string) error {
	return confirmCheckboxOrRadio(s, true)
}

func shouldNotBeSelected(s *steputils.StepUtils) func(string) error {
	return confirmCheckboxOrRadio(s, false)
}

func confirmCheckboxOrRadio(s *steputils.StepUtils, expectChecked bool) func(string) error {
	return func(field string) error {
		var (
			matches []*agouti.Selection
			errMsg  = fmt.Sprintf("error encountered while verifying selected status of [%s]: ", field) + "%s"
			err     error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's at least 1
		fieldCnt := len(matches)
		if fieldCnt < 1 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// ensure its a checkbox or radio
		if fieldType, err := s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, fmt.Sprintf("error encountered determining field type: %s", err))
		} else if !strings.EqualFold(fieldType, "checkbox") && !strings.EqualFold(fieldType, "radio") {
			return fmt.Errorf(errMsg, "[%s] must be of type [checkbox] or [radio] but is [%s]", field, fieldType)
		}

		// determine checked status
		checkboxElms, _ := matches[0].Elements()
		isChecked, _ := checkboxElms[0].IsSelected()

		// verify checked status
		if (isChecked && expectChecked) || !(isChecked || expectChecked) {
			return nil
		}

		// if it failed, produce an appropriate message
		shouldMsg := "should "
		if expectChecked {
			shouldMsg += "be selected but is not"
		} else {
			shouldMsg += "not be selected but is"
		}

		return fmt.Errorf(errMsg, shouldMsg)
	}
}

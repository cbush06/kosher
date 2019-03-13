package websteps

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func iFillInTheFollowing(s *steputils.StepUtils) func(*gherkin.DataTable) error {
	return func(fields *gherkin.DataTable) error {
		fillInFieldFunc := iFillInFieldWith(s)

		for _, row := range fields.Rows {
			var (
				field = strings.TrimSpace(row.Cells[0].Value)
				value = row.Cells[1].Value
			)

			if err := fillInFieldFunc(field, value); err != nil {
				return err
			}
		}
		return nil
	}
}

func iFillInFieldWith(s *steputils.StepUtils) func(string, string) error {
	return func(field string, value string) error {
		var (
			matches     []*agouti.Selection
			matchCnt    int
			fieldType   string
			shouldCheck bool
			errMsg      = "error encountered while filling in multiple fields: %s"
			err         error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// based on the field type, we apply the actual value
		if s.IsTextBased(field, matches[0]) {
			matches[0].Fill(value)
		} else {
			matchCnt = len(matches)
			switch fieldType {
			case "radio":
				for i := 0; i < matchCnt; i++ {
					nextRadioEls, _ := matches[i].Elements()
					nextRadioValue, _ := nextRadioEls[0].GetAttribute("value")
					if strings.EqualFold(nextRadioValue, strings.TrimSpace(value)) {
						matches[i].Click()
						return nil
					}
				}
			case "checkbox":
				shouldCheck, err = strconv.ParseBool(strings.TrimSpace(value))
				if shouldCheck && err == nil {
					matches[0].Check()
					return nil
				}
			case "select":
				if err = matches[0].Select(value); err != nil {
					return fmt.Errorf(errMsg, err)
				}
			default:
				return fmt.Errorf("unrecognized field type [%s]", fieldType)
			}
		}
		return nil
	}
}

func iKeyIn(s *steputils.StepUtils) func(string, string) error {
	return func(value string, field string) error {
		var (
			matches   []*agouti.Selection
			fieldType string
			err       error
			errMsg    = fmt.Sprintf("error encountered while keying [%s] into [%s]: ", value, field) + "%s"
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		if s.IsTextBased(field, matches[0]) {
			// fill it in
			if err = matches[0].SendKeys(value); err != nil {
				return fmt.Errorf(errMsg, err)
			}
		} else {
			return fmt.Errorf(errMsg, fmt.Sprintf("field is of type [%s] but must be a text-based field", fieldType))
		}

		return nil
	}
}

func iSelectFrom(s *steputils.StepUtils) func(string, string) error {
	return func(value string, field string) error {
		var (
			matches   []*agouti.Selection
			fieldType string
			errMsg    = fmt.Sprintf("error encountered while selecting [%s] from [%s]: ", value, field) + "%s"
			err       error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure it's a select
		if fieldType != "select" {
			return fmt.Errorf(errMsg, "field is of type [%s] but must be type [select]", fieldType)
		}

		// select the value
		if err = matches[0].Select(value); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		return nil
	}
}

func iCheck(s *steputils.StepUtils) func(string) error {
	return iCheckUncheck(s, true)
}

func iUncheck(s *steputils.StepUtils) func(string) error {
	return iCheckUncheck(s, false)
}

func iCheckUncheck(s *steputils.StepUtils, checked bool) func(string) error {
	return func(field string) error {
		var (
			matches   []*agouti.Selection
			fieldType string
			errMsg    = fmt.Sprintf("error encountered while unchecking [%s]: ", field) + "%s"
			err       error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure it's a select
		if fieldType != "checkbox" {
			return fmt.Errorf(errMsg, "field is of type [%s] but must be type [checkbox]", fieldType)
		}

		// set the value
		if checked {
			if err = matches[0].Check(); err != nil {
				return fmt.Errorf(errMsg, err)
			}
		} else {
			if err = matches[0].Uncheck(); err != nil {
				return fmt.Errorf(errMsg, err)
			}
		}

		return nil
	}
}

func iSelectTheFollowingValues(s *steputils.StepUtils) func(string, *gherkin.DataTable) error {
	return iSelectUnselectTheFollowingValues(s, true)
}

func iUnselectTheFollowingValues(s *steputils.StepUtils) func(string, *gherkin.DataTable) error {
	return iSelectUnselectTheFollowingValues(s, false)
}

func iSelectUnselectTheFollowingValues(s *steputils.StepUtils, doSelect bool) func(string, *gherkin.DataTable) error {
	return func(field string, values *gherkin.DataTable) error {
		var (
			matches      []*agouti.Selection
			fieldType    string
			errMsg       = fmt.Sprintf("error encountered while selecting/unselecting multiple values from [%s]: ", field) + "%s"
			err          error
			valuesLookup []string
			isSelected   bool
		)

		// build a sorted array of values
		for _, nextRow := range values.Rows {
			valuesLookup = append(valuesLookup, strings.TrimSpace(nextRow.Cells[0].Value))
		}
		sort.Strings(valuesLookup)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// only handle multi-selects
		if fieldType == "select" {
			selectElms, _ := matches[0].Elements()
			isMultiple, _ := selectElms[0].GetAttribute("multiple")

			if !strings.EqualFold(isMultiple, "true") && !strings.EqualFold(isMultiple, "multiple") {
				return fmt.Errorf(errMsg, "cannot select/unselect multiple values on a single select list")
			}

			for _, nextValue := range valuesLookup {
				nextOption := matches[0].AllByXPath(`option[normalize-space(text())="` + nextValue + `"]`)
				if isSelected, err = nextOption.Selected(); err != nil {
					return fmt.Errorf(errMsg, err)
				} else if (doSelect && !isSelected) || (!doSelect && isSelected) {
					nextOption.Click()
				}
			}
		} else {
			return fmt.Errorf(errMsg, "field is of type [%s] but must be type [select,checkbox]", fieldType)
		}

		return nil
	}
}

func iChoose(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		var (
			matches   []*agouti.Selection
			fieldType string
			errMsg    = fmt.Sprintf("error encountered while choosing radio button [%s]: ", field) + "%s"
			err       error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches[0]); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// if it's a radio, click it
		if fieldType == "radio" {
			matches[0].Click()
			return nil
		}

		return fmt.Errorf(errMsg, fmt.Sprintf("must be of type [radio] but is of type [%s]", fieldType))
	}
}

func iPress(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		return iClickIdx(s, field, 0)
	}
}

func iClickTheButtonLink(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		return iClickIdx(s, field, 0)
	}
}

func iPressNth(s *steputils.StepUtils) func(string, string) error {
	return func(nth string, field string) error {
		var (
			matches    []*agouti.Selection
			errMsg     = fmt.Sprintf("error encountered while pressing button/link [%s]: ", field) + "%s"
			nthNumeric int
			err        error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's enough of these to satisfy the `nth` argument
		fieldCnt := len(matches)
		if fieldCnt < 1 {
			return fmt.Errorf(errMsg, "no matching buttons/links found")
		}

		// if something other than an integer was specified
		if nthNumeric, err = strconv.Atoi(nth); err != nil {
			// if "first" or "last" was specified
			if strings.EqualFold(nth, "first") {
				return iClickIdx(s, field, 0)
			}
			if strings.EqualFold(nth, "last") {
				return iClickIdx(s, field, fieldCnt-1)
			}

			re := regexp.MustCompile(`\d+`)

			var match string
			if match = re.FindString(nth); len(match) < 1 {
				return fmt.Errorf("no valid position specified for button")
			}
			nthNumeric, _ = strconv.Atoi(match)
		}

		if nthNumeric > fieldCnt {
			return fmt.Errorf("requested click on button [%d], but only [%d] buttons found", nthNumeric, fieldCnt)
		}
		return iClickIdx(s, field, nthNumeric-1)

	}
}

func iClickIdx(s *steputils.StepUtils, field string, btnNumber int) error {
	var (
		matches []*agouti.Selection
		errMsg  = fmt.Sprintf("error encountered while clicking button/link [%s]: ", field) + "%s"
		err     error
		count   int
	)

	// try to find the field(s) specified
	if matches, err = s.ResolveSelector(field); err != nil {
		return fmt.Errorf(errMsg, err)
	}

	// confirm a match was found
	if count = len(matches); count < (btnNumber - 1) {
		return fmt.Errorf(errMsg, fmt.Sprintf("instance [%d] not found", btnNumber))
	}

	matches[btnNumber].Click()

	return nil
}

func iUnfocusBlur(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		var (
			matches []*agouti.Selection
			errMsg  = fmt.Sprintf("error encountered while unfocusing/blurring [%s]: ", field) + "%s"
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

		// unfocus the field
		fieldElms, _ := matches[0].Elements()
		fieldID, _ := fieldElms[0].GetAttribute("id")
		fieldName, _ := fieldElms[0].GetAttribute("name")
		var fieldSelector string

		if len(fieldID) > 0 {
			fieldSelector = "#" + fieldID
		} else if len(fieldName) > 0 {
			fieldSelector = "*[name='" + fieldName + "']"
		}

		if len(fieldSelector) > 0 {
			// if a fieldSelector could be derived, use document.getElement
			err = s.Page.RunScript(`
				const selectedEl = document.querySelector( selector );
				if( selectedEl != null ) {
					selectedEl.blur();
				}
			`, map[string]interface{}{"selector": fieldSelector}, nil)
		} else {
			// if no fieldSelector was derived, try sending a TAB key
			err = matches[0].SendKeys("\t")
		}

		if err != nil {
			return fmt.Errorf(errMsg, err)
		}

		return nil
	}
}

func iEnterTodaysDateIn(s *steputils.StepUtils) func(string) error {
	today := s.FormatDate(time.Now())
	fillInFieldFunc := iFillInFieldWith(s)
	return func(field string) error {
		return fillInFieldFunc(field, today)
	}
}

package websteps

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sclevine/agouti"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/cbush06/kosher/steps/steputils"
)

func iShouldSee(s *steputils.StepUtils) func(string) error {
	return confirmSee(s, true)
}

func iShouldNotSee(s *steputils.StepUtils) func(string) error {
	return confirmSee(s, false)
}

func iShouldSeeAllOfTheTexts(s *steputils.StepUtils) func(*gherkin.DataTable) error {
	confirmSeeFunc := confirmSee(s, true)

	return func(textsTable *gherkin.DataTable) error {
		for _, nextRow := range textsTable.Rows {
			if err := confirmSeeFunc(nextRow.Cells[0].Value); err != nil {
				return err
			}
		}
		return nil
	}
}

func confirmSee(s *steputils.StepUtils, shouldSee bool) func(string) error {
	return func(text string) error {
		var (
			count int
			err   error
		)

		matches := s.Page.AllByXPath(fmt.Sprintf(`//*[text()[contains(., "%s")]]`, text))
		if count, err = matches.Count(); err != nil {
			return fmt.Errorf(`error encountered searching page for text: %s`, err)
		}

		if shouldSee && (count < 1) {
			return fmt.Errorf(`expected to find [%s] but did not`, text)
		} else if !shouldSee && (count > 0) {
			return fmt.Errorf(`expected NOT to find [%s] but found it [%d] times`, text, count)
		}

		return nil
	}
}

func iShouldSeeButtonLink(s *steputils.StepUtils) func(string) error {
	return confirmSeeButtonLink(s, true)
}

func iShouldNotSeeButtonLink(s *steputils.StepUtils) func(string) error {
	return confirmSeeButtonLink(s, false)
}

func confirmSeeButtonLink(s *steputils.StepUtils, shouldSee bool) func(string) error {
	return func(text string) error {
		var (
			matches    *agouti.MultiSelection
			err        error
			errMsg     = fmt.Sprintf(`error encountered while searching for [%s] button/link: `, text) + "%s"
			found      bool
			matchCount int
		)

		// resolve selector...not getting an error means a match was found
		if matches, err = s.ResolveSelector(text); err == nil {
			// confirm one of matches is button or link
			count, _ := matches.Count()
			for i := 0; i < count; i++ {
				if fieldType, err2 := s.GetFieldType("text", matches.At(i)); err2 != nil {
					return fmt.Errorf(errMsg, err2)
				} else if strings.EqualFold(fieldType, "a") || strings.EqualFold(fieldType, "button") {
					found = true
					matchCount++
				}
			}
		}

		if (shouldSee && found) || !(shouldSee || found) {
			return nil
		} else if found {
			return fmt.Errorf(`expected to not find elements for identifier [%s] but found [%d]`, text, matchCount)
		} else {
			return fmt.Errorf(`no element of type link/button found for identifier [%s]`, text)
		}
	}
}

func theNthInstanceOfShouldBeEnabled(s *steputils.StepUtils) func(string, string) error {
	return confirmDisabled(s, false)
}

func theNthInstanceOfShouldBeDisabled(s *steputils.StepUtils) func(string, string) error {
	return confirmDisabled(s, true)
}

func confirmDisabled(s *steputils.StepUtils, shouldBeDisabled bool) func(string, string) error {
	return func(nth string, field string) error {
		var (
			matches    *agouti.MultiSelection
			errMsg     = fmt.Sprintf("error encountered while confirming [%s] instance of [%s] is disabled: ", nth, field) + "%s"
			nthNumeric int
			err        error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's enough of these to satisfy the `nth` argument
		fieldCnt, _ := matches.Count()
		if fieldCnt < 0 {
			return fmt.Errorf(errMsg, "no matching elements found")
		}

		// if something other than an integer was specified
		if nthNumeric, err = strconv.Atoi(nth); err != nil {
			if strings.EqualFold(nth, "first") {
				// if "first" was specified
				nthNumeric = 0
			} else if strings.EqualFold(nth, "last") {
				// if "last" was specified
				nthNumeric = fieldCnt - 1
			} else {
				// if something else
				re := regexp.MustCompile(`\d+`)

				var match string
				if match = re.FindString(nth); len(match) < 1 {
					return fmt.Errorf("no valid position specified for button")
				}
				nthNumeric, _ = strconv.Atoi(match)
			}
		}

		if nthNumeric > fieldCnt {
			return fmt.Errorf("specified element [%d], but only [%d] found", nthNumeric, fieldCnt)
		}

		// verify it's disabled
		fieldElms, _ := matches.At(nthNumeric).Elements()
		enabled, _ := fieldElms[0].IsEnabled()

		if shouldBeDisabled && enabled {
			return fmt.Errorf("the [%s] instance of [%s] is not disabled", nth, field)
		} else if !shouldBeDisabled && !enabled {
			return fmt.Errorf("the [%s] instance of [%s] is not enabled", nth, field)
		}
		return nil
	}
}

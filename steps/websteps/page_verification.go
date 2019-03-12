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
			matchCount   int
			visibleCount int
			isShown      bool
			width        int
			height       int
			err          error
		)

		matches := s.Page.AllByXPath(fmt.Sprintf(`//*[text()[contains(., "%s") and not(ancestor::*[contains(@style, 'display: none')])]]`, text))
		if matchCount, err = matches.Count(); err != nil {
			return fmt.Errorf(`error encountered searching page for text: %s`, err)
		}

		if shouldSee && (matchCount < 1) {
			return fmt.Errorf(`expected to find [%s] but did not`, text)
		}

		// determine how many of these elements are visible
		for i := 0; i < matchCount; i++ {
			els, _ := matches.At(i).Elements()
			if isShown, err = els[0].IsDisplayed(); err != nil {
				return fmt.Errorf(`error encountered determining if matched text is visible: %s`, err)
			}
			if width, height, err = els[0].GetSize(); err != nil {
				return fmt.Errorf(`error encountered determining size of element containing matched text: %s`, err)
			}
			if isShown && (width > 0 || height > 0) {
				visibleCount++
			}
		}

		// show error message if applicable
		if shouldSee && visibleCount < 1 {
			return fmt.Errorf(`expected TO find [%s] but did not`, text)
		} else if !shouldSee && visibleCount > 0 {
			return fmt.Errorf(`expected NOT to find [%s] but found it [%d] times`, text, visibleCount)
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
			matches           []*agouti.Selection
			buttonLinkMatches []*agouti.Selection
			err               error
			visibleCount      int
			isShown           bool
			width             int
			height            int
			errMsg            = fmt.Sprintf(`error encountered while searching for [%s] button/link: `, text) + "%s"
		)

		// resolve selector...not getting an error means a match was found
		if matches, err = s.ResolveSelector(text); err == nil {
			// confirm one of matches is button or link
			count := len(matches)
			for i := 0; i < count; i++ {
				if fieldType, err2 := s.GetFieldType(text, matches[i]); err2 != nil {
					return fmt.Errorf(errMsg, err2)
				} else if strings.EqualFold(fieldType, "a") || strings.EqualFold(fieldType, "button") {
					buttonLinkMatches = append(buttonLinkMatches, matches[i])
				}
			}
		}

		// determine how many of these elements are visible
		for _, match := range buttonLinkMatches {
			els, _ := match.Elements()
			if isShown, err = els[0].IsDisplayed(); err != nil {
				return fmt.Errorf(`error encountered determining if matched button/link is visible: %s`, err)
			}
			if width, height, err = els[0].GetSize(); err != nil {
				return fmt.Errorf(`error encountered determining size of matched button/link: %s`, err)
			}
			if isShown && (width > 0 || height > 0) {
				visibleCount++
			}
		}

		// show error message if applicable
		if shouldSee && visibleCount < 1 {
			return fmt.Errorf(`expected TO find button/link [%s] but did not`, text)
		} else if !shouldSee && visibleCount > 0 {
			return fmt.Errorf(`expected NOT to find button/link [%s] but found it [%d] times`, text, visibleCount)
		}

		return nil
	}
}

func shouldBeDisabled(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		return confirmDisabled(s, true)("first", field)
	}
}

func shouldBeEnabled(s *steputils.StepUtils) func(string) error {
	return func(field string) error {
		return confirmDisabled(s, false)("first", field)
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
			matches    []*agouti.Selection
			errMsg     = fmt.Sprintf("error encountered while confirming [%s] instance of [%s] is disabled: ", nth, field) + "%s"
			nthNumeric int
			err        error
		)

		// try to find the field(s) specified
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// ensure there's enough of these to satisfy the `nth` argument
		fieldCnt := len(matches)
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
		fieldElms, _ := matches[nthNumeric].Elements()
		enabled, _ := fieldElms[0].IsEnabled()

		if shouldBeDisabled && enabled {
			return fmt.Errorf("the [%s] instance of [%s] is not disabled", nth, field)
		} else if !shouldBeDisabled && !enabled {
			return fmt.Errorf("the [%s] instance of [%s] is not enabled", nth, field)
		}
		return nil
	}
}

func confirmSeeOfType(s *steputils.StepUtils) func(string, string, string) error {
	return confirmSeeNotSeeOfType(s, true)
}

func confirmNotSeeOfType(s *steputils.StepUtils) func(string, string, string) error {
	return confirmSeeNotSeeOfType(s, false)
}

func confirmSeeNotSeeOfType(s *steputils.StepUtils, shouldSee bool) func(string, string, string) error {
	return func(tag string, attr string, attrVal string) error {
		var (
			matches      *agouti.MultiSelection
			matchCount   int
			visibleCount int
			isShown      bool
			width        int
			height       int
			err          error
			errMsg       = fmt.Sprintf(`error encountered searching page for [%s] with [%s] of [%s]: `, tag, attr, attrVal) + "%s"
		)

		switch attr {
		case "id":
			matches = s.Page.AllByXPath(fmt.Sprintf(`//%s[@id="%s"]`, strings.ToLower(tag), attrVal))
		case "name":
			matches = s.Page.AllByXPath(fmt.Sprintf(`//%s[@name="%s"]`, strings.ToLower(tag), attrVal))
		default:
			return fmt.Errorf(errMsg, `expected attribute to be either [id] or [name]`)
		}

		if matchCount, err = matches.Count(); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		if shouldSee && (matchCount < 1) {
			return fmt.Errorf(errMsg, "expected TO find match but did not")
		}

		// determine how many of these elements are visible
		for i := 0; i < matchCount; i++ {
			els, _ := matches.At(i).Elements()

			if len(els) < 1 {
				continue
			}

			if isShown, err = els[0].IsDisplayed(); err != nil {
				return fmt.Errorf(errMsg, fmt.Sprintf(`error encountered determining if matched element is visible: %s`, err))
			}
			if width, height, err = els[0].GetSize(); err != nil {
				return fmt.Errorf(errMsg, fmt.Sprintf(`error encountered determining size of matched element: %s`, err))
			}
			if isShown && (width > 0 || height > 0) {
				visibleCount++
			}
		}

		// show error message if applicable
		if shouldSee && visibleCount < 1 {
			return fmt.Errorf(`expected TO find [%s] with [%s] of [%s], but did not`, tag, attr, attrVal)
		} else if !shouldSee && visibleCount > 0 {
			return fmt.Errorf(`expected NOT to find [%s] with [%s] of [%s] but found it [%d] times`, tag, attr, attrVal, visibleCount)
		}

		return nil
	}
}

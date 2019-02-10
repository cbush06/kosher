package websteps

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func iFillInTheFollowing(s *steputils.StepUtils, p *agouti.Page) func(*gherkin.DataTable) error {
	return func(fields *gherkin.DataTable) error {
		fillInFieldFunc := iFillInFieldWith(s, p)

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

func iFillInFieldWith(s *steputils.StepUtils, p *agouti.Page) func(string, string) error {
	return func(field string, value string) error {
		var (
			matches   *agouti.MultiSelection
			matchCnt  int
			fieldType string
			errMsg    = "error encountered while filling in multiple fields: %s"
			err       error
		)

		// check if the user provided a label or a selector
		if matches, err = s.ResolveSelector(field); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// determine the field's type
		if fieldType, err = s.GetFieldType(field, matches.At(0)); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// based on the field type, we apply the actual value
		matchCnt, _ = matches.Count()
		switch fieldType {
		case "text":
			matches.Fill(value)
		case "radio":
			for i := 0; i < matchCnt; i++ {
				nextRadioEls, _ := matches.At(i).Elements()
				nextRadioValue, _ := nextRadioEls[0].GetAttribute("value")
				if strings.EqualFold(nextRadioValue, strings.TrimSpace(value)) {
					matches.At(i).Click()
					return nil
				}
			}
		}
		return nil
	}
}

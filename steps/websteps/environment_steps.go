package websteps

import (
	"fmt"

	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

func iMaximizeTheWindow(s *steputils.StepUtils, p *agouti.Page) func() error {
	return func() error {
		if err := p.Size(s.GetMaxWindowSize()); err != nil {
			return fmt.Errorf("error encountered resizing window: %s", err)
		}
		return nil
	}
}

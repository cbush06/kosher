package websteps

import (
	"fmt"

	"github.com/cbush06/kosher/steps/steputils"
)

// I am on the "<name>" page
func iAmOnThePage(s *steputils.StepUtils) func(string) error {
	return func(pageName string) error {
		var (
			url string
			err error
		)

		if url, err = s.ResolvePage(pageName); err != nil {
			return fmt.Errorf("error encountered while resolving page URL: %s", err)
		} else if len(url) < 1 {
			return fmt.Errorf("no page found for name [%s]", pageName)
		}

		if err = s.Page.Navigate(url); err != nil {
			return fmt.Errorf("failed to load page [%s]: %s", pageName, err)
		}
		return nil
	}
}

// I go to the "<name>" page
func iGoToThePage(s *steputils.StepUtils) func(string) error {
	return iAmOnThePage(s)
}

// I follow "<link label>"
func iFollow(s *steputils.StepUtils) func(string) error {
	return func(linkLabel string) error {
		if err := s.Page.FirstByLink(linkLabel).Click(); err != nil {
			return fmt.Errorf("failed to click link with label [%s]: %s", linkLabel, err)
		}
		return nil
	}
}

// I switch to frame 2
func iSwitchToFrameN(s *steputils.StepUtils) func(int) error {
	return func(frameNumber int) error {
		if err := s.Page.First("frame:nth-of-type(" + string(frameNumber) + ")").SwitchToFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to FRAME(%d) in page: %s", frameNumber, err)
		}
		return nil
	}
}

// I switch to iframe 2
func iSwitchToIFrameN(s *steputils.StepUtils) func(int) error {
	return func(frameNumber int) error {
		if err := s.Page.All("iframe").At(frameNumber).SwitchToFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to IFRAME(%d) in page: %s", frameNumber, err)
		}
		return nil
	}
}

// I switch to the root frame
func iSwitchToTheRootFrame(s *steputils.StepUtils) func() error {
	return func() error {
		if err := s.Page.SwitchToRootFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to the root frame: %s", err)
		}
		return nil
	}
}

// I accept the popup
func iAcceptThePopup(s *steputils.StepUtils) func() error {
	return func() error {
		if err := s.Page.ConfirmPopup(); err != nil {
			return fmt.Errorf("error encountered while accepting popup: %s", err)
		}
		return nil
	}
}

// I decline the popup
func iDeclineThePopup(s *steputils.StepUtils) func() error {
	return func() error {
		if err := s.Page.CancelPopup(); err != nil {
			return fmt.Errorf("error encountered while declining popup: %s", err)
		}
		return nil
	}
}

// I enter "Some Text" in the popup
func iEnterInThePopup(s *steputils.StepUtils) func(string) error {
	return func(text string) error {
		if err := s.Page.EnterPopupText(text); err != nil {
			return fmt.Errorf("error encountered while entering [%s] in the popup: %s", text, err)
		}
		return nil
	}
}

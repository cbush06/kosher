package websteps

import (
	"fmt"
	"net/url"

	"github.com/cbush06/kosher/steps/steputils"
)

// I am on the "<name>" page
func iAmOnThePage(s *steputils.StepUtils) func(string) error {
	return func(pageName string) error {
		var (
			pageURL string
			err     error
		)

		if pageURL, err = s.ResolvePage(pageName); err != nil {
			return fmt.Errorf("error encountered while resolving page URL: %s", err)
		} else if len(pageURL) < 1 {
			return fmt.Errorf("no page found for name [%s]", pageName)
		}

		unescapedURL, _ := url.PathUnescape(pageURL)
		if err = s.Page.Navigate(unescapedURL); err != nil {
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
			return fmt.Errorf("error encountered while switching to FRAME(%d) in page: %s", frameNumber-1, err)
		}
		return nil
	}
}

// I switch to iframe 2
func iSwitchToIFrameN(s *steputils.StepUtils) func(int) error {
	return func(frameNumber int) error {
		if err := s.Page.All("iframe").At(frameNumber).SwitchToFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to IFRAME(%d) in page: %s", frameNumber-1, err)
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
		var (
			interpolatedText string
			err              error
			errMsg           = "error encountered while entering [%s] in the popup: %s"
		)

		if interpolatedText, err = s.ReplaceVariables(text); err != nil {
			return fmt.Errorf(errMsg, text, err)
		}

		if err := s.Page.EnterPopupText(interpolatedText); err != nil {
			return fmt.Errorf(errMsg, text, err)
		}
		return nil
	}
}

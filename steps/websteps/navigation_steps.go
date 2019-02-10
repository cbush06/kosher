package websteps

import (
	"fmt"

	"github.com/cbush06/kosher/config"
	"github.com/sclevine/agouti"
)

// I am on the "<name>" page
func iAmOnThePage(s *config.Settings, p *agouti.Page) func(pageName string) error {
	return func(pageName string) error {
		url := s.Pages.GetString(pageName)
		if len(url) < 1 {
			return fmt.Errorf("no page found for name [%s]", pageName)
		}

		if err := p.Navigate(url); err != nil {
			return fmt.Errorf("failed to load page [%s]: %s", pageName, err)
		}
		return nil
	}
}

// I go to the "<name>" page
func iGoToThePage(s *config.Settings, p *agouti.Page) func(pageName string) error {
	return iAmOnThePage(s, p)
}

// I follow "<link label>"
func iFollow(p *agouti.Page) func(linkLabel string) error {
	return func(linkLabel string) error {
		if err := p.FirstByLink(linkLabel).Click(); err != nil {
			return fmt.Errorf("failed to click link with label [%s]: %s", linkLabel, err)
		}
		return nil
	}
}

// I switch to frame 2
func iSwitchToFrameN(p *agouti.Page) func(int) error {
	return func(frameNumber int) error {
		if err := p.First("frame:nth-of-type(" + string(frameNumber) + ")").SwitchToFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to FRAME(%d) in page: %s", frameNumber, err)
		}
		return nil
	}
}

// I switch to iframe 2
func iSwitchToIFrameN(p *agouti.Page) func(int) error {
	return func(frameNumber int) error {
		if err := p.All("iframe").At(frameNumber).SwitchToFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to IFRAME(%d) in page: %s", frameNumber, err)
		}
		return nil
	}
}

// I switch to the root frame
func iSwitchToTheRootFrame(p *agouti.Page) func() error {
	return func() error {
		if err := p.SwitchToRootFrame(); err != nil {
			return fmt.Errorf("error encountered while switching to the root frame: %s", err)
		}
		return nil
	}
}

// I accept the popup
func iAcceptThePopup(p *agouti.Page) func() error {
	return func() error {
		if err := p.ConfirmPopup(); err != nil {
			return fmt.Errorf("error encountered while accepting popup: %s", err)
		}
		return nil
	}
}

// I decline the popup
func iDeclineThePopup(p *agouti.Page) func() error {
	return func() error {
		if err := p.CancelPopup(); err != nil {
			return fmt.Errorf("error encountered while declining popup: %s", err)
		}
		return nil
	}
}

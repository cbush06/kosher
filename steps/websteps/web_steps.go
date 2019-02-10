package websteps

import (
	"github.com/DATA-DOG/godog"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti"
)

// BuildGoDogSuite loads the provided godog suite with step definitions for the provided page
func BuildGoDogSuite(settings *config.Settings, page *agouti.Page, suite *godog.Suite) {
	utils := steputils.NewStepUtils(settings, page)

	suite.Step(`^(?:|I )am on the "(.*)" page$`, iAmOnThePage(settings, page))
	suite.Step(`^(?:|I )go to the "(.*)" page$`, iGoToThePage(settings, page))
	suite.Step(`^(?:|I )follow "([^"]*)"$`, iFollow(page))
	suite.Step(`^(?:|I )maximize the window$`, iMaximizeTheWindow(utils, page))
	suite.Step(`^(?:|I )switch to frame (\d+)$`, iSwitchToFrameN(page))
	suite.Step(`^(?:|I )switch to iframe (\d+)$`, iSwitchToIFrameN(page))
	suite.Step(`^(?:|I )switch to the root frame$`, iSwitchToTheRootFrame(page))
	suite.Step(`^(?:|I )accept the popup$`, iAcceptThePopup(page))
	suite.Step(`^(?:|I )decline the popup$`, iDeclineThePopup(page))
}

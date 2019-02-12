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

	// environment steps
	suite.Step(`^(?:|I )wait (\d+) seconds$`, iWaitSeconds())
	suite.Step(`^(?:|I )maximize the window$`, iMaximizeTheWindow(utils))

	// navigation steps
	suite.Step(`^(?:|I )am on the "([^"]*)" page$`, iAmOnThePage(utils))
	suite.Step(`^(?:|I )go to the "([^"]*)" page$`, iGoToThePage(utils))
	suite.Step(`^(?:|I )follow "([^"]*)"$`, iFollow(utils))
	suite.Step(`^(?:|I )switch to frame (\d+)$`, iSwitchToFrameN(utils))
	suite.Step(`^(?:|I )switch to iframe (\d+)$`, iSwitchToIFrameN(utils))
	suite.Step(`^(?:|I )switch to the root frame$`, iSwitchToTheRootFrame(utils))
	suite.Step(`^(?:|I )accept the popup$`, iAcceptThePopup(utils))
	suite.Step(`^(?:|I )decline the popup$`, iDeclineThePopup(utils))

	// form steps
	suite.Step(`^(?:|I )fill in "([^"]*)" with "([^"]*)"$`, iFillInFieldWith(utils))
	suite.Step(`^(?:|I )fill in the following:$`, iFillInTheFollowing(utils))
	suite.Step(`^(?:|I )key in "([^"]*)" in the "([^"]*)" field$`, iKeyInTheField(utils))
	suite.Step(`^(?:|I )key "([^"]*)" in "([^"]*)"$`, iKeyIn(utils))
	suite.Step(`^(?:|I )select "([^"]*)" from "([^"]*)"$`, iSelectFrom(utils))
	suite.Step(`^(?:|I )check "([^"]*)"$`, iCheck(utils))
	suite.Step(`^(?:|I )uncheck "([^"]*)"$`, iUncheck(utils))
	suite.Step(`^(?:|I )select (?:the )following values from "([^"]*)":$`, iSelectTheFollowingValues(utils))
	suite.Step(`^(?:|I )unselect (?:the )following values from "([^"]*)":$`, iUnselectTheFollowingValues(utils))
	suite.Step(`^(?:|I )choose "([^"]*)"$`, iChoose(utils))
	suite.Step(`^(?:|I )press "([^"]*)"$`, iPress(utils))
	suite.Step(`^(?:|I )(?:press|click) the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)"$`, iPressNth(utils))
	suite.Step(`^(?:|I )click the "([^"]*)" (?:button|link)$`, iClickTheButtonLink(utils))
	suite.Step(`^(?:|I )(?:unfocus|blur) "([^"]*)"$`, iUnfocusBlur(utils))
	suite.Step(`^(?:|I )hover over "([^"]*)"$`, iHoverOver(utils))
	suite.Step(`^(?:|I )enter today's date in "([^"]*)"`, iEnterTodaysDateIn(utils))

	// validation steps
	suite.Step(`^(?:|I )verify "([^"]*)" has today's date$`, iVerifyHasTodaysDate(utils))
	suite.Step(`^"([^"]*)" should contain "([^"]*)"$`, shouldContain(utils))
	suite.Step(`^"([^"]*)" should not contain "([^"]*)"$`, shouldNotContain(utils))
}

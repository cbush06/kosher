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
	suite.Step(`^(?:|I )wait (\d+) second(?:|s)$`, iWaitSeconds())
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
	suite.Step(`^(?:|I )key (?:|in )"([^"]*)" in (?:|the )"([^"]*)"(?:| field)$`, iKeyIn(utils))
	suite.Step(`^(?:|I )select "([^"]*)" from "([^"]*)"$`, iSelectFrom(utils))
	suite.Step(`^(?:|I )check "([^"]*)"$`, iCheck(utils))
	suite.Step(`^(?:|I )uncheck "([^"]*)"$`, iUncheck(utils))
	suite.Step(`^(?:|I )select (?:|the )following values from "([^"]*)":$`, iSelectTheFollowingValues(utils))
	suite.Step(`^(?:|I )unselect (?:|the )following values from "([^"]*)":$`, iUnselectTheFollowingValues(utils))
	suite.Step(`^(?:|I )choose (?:|the )"([^"]*)" radio$`, iChoose(utils))
	suite.Step(`^(?:|I )press "([^"]*)"$`, iPress(utils))
	suite.Step(`^(?:|I )(?:press|click) the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)"$`, iPressNth(utils))
	suite.Step(`^(?:|I )click the "([^"]*)" (?:button|link)$`, iClickTheButtonLink(utils))
	suite.Step(`^(?:|I )(?:unfocus|blur) "([^"]*)"$`, iUnfocusBlur(utils))
	suite.Step(`^(?:|I )hover over "([^"]*)"$`, iHoverOver(utils))
	suite.Step(`^(?:|I )enter today's date in "([^"]*)"`, iEnterTodaysDateIn(utils))

	// form verification steps
	suite.Step(`^(?:|I )verify "([^"]*)" has today's date$`, iVerifyHasTodaysDate(utils))
	suite.Step(`^"([^"]*)" should contain "([^"]*)"$`, shouldContain(utils))
	suite.Step(`^"([^"]*)" should not contain "([^"]*)"$`, shouldNotContain(utils))
	suite.Step(`^"([^"]*)" should have (?:|the )following options:$`, shouldHaveTheFollowingOptions(utils))
	suite.Step(`^"([^"]*)" should have (?:|the )following options selected:$`, shouldHaveTheFollowingOptionsSelected(utils))
	suite.Step(`^"([^"]*)" should not have (?:|the )following options selected:$`, shouldNotHaveTheFollowingOptionsSelected(utils))
	suite.Step(`^the "([^"]*)" checkbox should be checked$`, shouldBeSelected(utils))
	suite.Step(`^the "([^"]*)" checkbox should not be checked$`, shouldNotBeSelected(utils))
	suite.Step(`^the "([^"]*)" radio should be selected$`, shouldBeSelected(utils))
	suite.Step(`^the "([^"]*)" radio should not be selected$`, shouldNotBeSelected(utils))

	// navigation verification steps
	suite.Step(`^I should be redirected to the "([^"]*)" page$`, iShouldBeRedirectedTo(utils))
	suite.Step(`^(?:|I )should be on the "([^"]*)" page$`, iShouldBeOn(utils))

	// page verification steps
	suite.Step(`^(?:|I )should see "([^"]*)"$`, iShouldSee(utils))
	suite.Step(`^(?:|I )should not see "([^"]*)"$`, iShouldNotSee(utils))
	suite.Step(`^(?:|I )should see all of the texts:$`, iShouldSeeAllOfTheTexts(utils))
	suite.Step(`^(?:|I )should see (?:|the )"([^"]*)"(?: button| link)$`, iShouldSeeButtonLink(utils))
	suite.Step(`^(?:|I )should not see (?:|the )"([^"]*)"(?: button| link)$`, iShouldNotSeeButtonLink(utils))
	suite.Step(`^the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)" should be disabled$`, theNthInstanceOfShouldBeDisabled(utils))
	suite.Step(`^the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)" should be enabled$`, theNthInstanceOfShouldBeEnabled(utils))
}

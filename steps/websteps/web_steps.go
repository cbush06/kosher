package websteps

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/spf13/afero"

	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/interfaces"
	"github.com/cbush06/kosher/steps/steputils"
)

// BuildGoDogSuite loads the provided godog suite with step definitions for the provided page
func BuildGoDogSuite(settings *config.Settings, page interfaces.PageService, suite *godog.Suite) {
	utils := steputils.NewStepUtils(settings, page)

	// browser steps
	suite.Step(`^(?:|I )wait (\d+) second(?:|s)$`, iWaitSeconds())
	suite.Step(`^(?:|I )maximize the window$`, iMaximizeTheWindow(utils))
	suite.Step(`^(?:|I )take a screenshot$`, iTakeAScreenshot(utils))
	suite.Step(`^(?:|I )switch to the (first|last|[0-9]+(?:th|st|rd|nd)) window$`, iSwitchToTheWindow(utils))
	suite.Step(`^(?:|I )reload the page$`, iReloadThePage(utils))
	suite.Step(`^(?:|I )key "([^"]*)"(?:| in the active element)$`, sendKeysToActiveElement(utils))

	// navigation steps
	suite.Step(`^(?:|I )am on the "([^"]*)" page$`, iAmOnThePage(utils))
	suite.Step(`^(?:|I )go to the "([^"]*)" page$`, iGoToThePage(utils))
	suite.Step(`^(?:|I )follow "([^"]*)"$`, iFollow(utils))
	suite.Step(`^(?:|I )switch to frame (\d+)$`, iSwitchToFrameN(utils))
	suite.Step(`^(?:|I )switch to iframe (\d+)$`, iSwitchToIFrameN(utils))
	suite.Step(`^(?:|I )switch to the root frame$`, iSwitchToTheRootFrame(utils))
	suite.Step(`^(?:|I )accept the popup$`, iAcceptThePopup(utils))
	suite.Step(`^(?:|I )decline the popup$`, iDeclineThePopup(utils))
	suite.Step(`^(?:|I )enter "([^"]*)" in the popup$`, iEnterInThePopup(utils))

	// form steps
	suite.Step(`^(?:|I )fill (?:|in )"([^"]*)" with "([^"]*)"$`, iFillInFieldWith(utils))
	suite.Step(`^(?:|I )fill (?:|in )"([^"]*)" with:$`, iFillInFieldWithMultiline(utils))
	suite.Step(`^(?:|I )fill in the following:$`, iFillInTheFollowing(utils))
	suite.Step(`^(?:|I )key (?:|in )"([^"]*)" in (?:|the )"([^"]*)"(?:| field)$`, iKeyIn(utils))
	suite.Step(`^(?:|I )press (?:|the )"([^"]*)" key (?:on|in) "([^"]*)"$`, iPressTheKeyOn(utils))
	suite.Step(`^(?:|I )select "([^"]*)" from "([^"]*)"$`, iSelectFrom(utils))
	suite.Step(`^(?:|I )check "([^"]*)"$`, iCheck(utils))
	suite.Step(`^(?:|I )uncheck "([^"]*)"$`, iUncheck(utils))
	suite.Step(`^(?:|I )select (?:|the )following values from "([^"]*)":$`, iSelectTheFollowingValues(utils))
	suite.Step(`^(?:|I )unselect (?:|the )following values from "([^"]*)":$`, iUnselectTheFollowingValues(utils))
	suite.Step(`^(?:|I )choose (?:|the )"([^"]*)" radio$`, iChoose(utils))
	suite.Step(`^(?:|I )(?:press|click) "([^"]*)"$`, iPress(utils))
	suite.Step(`^(?:|I )(?:press|click) the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)"$`, iPressNth(utils))
	suite.Step(`^(?:|I )(?:press|click) the "([^"]*)" (?:button|link)$`, iClickTheButtonLink(utils))
	suite.Step(`^(?:|I )(?:unfocus|blur) "([^"]*)"$`, iUnfocusBlur(utils))
	suite.Step(`^(?:|I )hover over "([^"]*)"$`, iHoverOver(utils))
	suite.Step(`^(?:|I )enter today's date in "([^"]*)"$`, iEnterTodaysDateIn(utils))

	// form verification steps
	suite.Step(`^"([^"]*)" should contain today's date$`, shouldContainTodaysDate(utils))
	suite.Step(`^(?:|I )verify "([^"]*)" has today's date$`, iVerifyHasTodaysDate(utils))
	suite.Step(`^"([^"]*)" should contain:$`, shouldContainMultiline(utils))
	suite.Step(`^"([^"]*)" should contain "([^"]*)"$`, shouldContain(utils))
	suite.Step(`^"([^"]*)" should not contain:$`, shouldNotContainMultiline(utils))
	suite.Step(`^"([^"]*)" should not contain "([^"]*)"$`, shouldNotContain(utils))
	suite.Step(`^"([^"]*)" should have (?:|the )following options:$`, shouldHaveTheFollowingOptions(utils))
	suite.Step(`^"([^"]*)" should have (?:|the )following options selected:$`, shouldHaveTheFollowingOptionsSelected(utils))
	suite.Step(`^"([^"]*)" should not have (?:|the )following options selected:$`, shouldNotHaveTheFollowingOptionsSelected(utils))
	suite.Step(`^the "([^"]*)" checkbox should be checked$`, shouldBeSelected(utils))
	suite.Step(`^the "([^"]*)" checkbox should not be checked$`, shouldNotBeSelected(utils))
	suite.Step(`^the "([^"]*)" radio should be selected$`, shouldBeSelected(utils))
	suite.Step(`^the "([^"]*)" radio should not be selected$`, shouldNotBeSelected(utils))

	// navigation verification steps
	suite.Step(`^(?:|I )should be redirected to the "([^"]*)" page$`, iShouldBeRedirectedTo(utils))
	suite.Step(`^(?:|I )should be on the "([^"]*)" page$`, iShouldBeOn(utils))
	suite.Step(`^(?:|I )should see (?:|the )popup (?:|text )"([^"]*)"$`, iShouldSeeThePopupText(utils))
	suite.Step(`^(?:|I )should not see (?:|the )popup (?:|text )"([^"]*)"$`, iShouldNotSeeThePopupText(utils))

	// page verification steps
	suite.Step(`^(?:|I )should see "([^"]*)"$`, iShouldSee(utils))
	suite.Step(`^(?:|I )should not see "([^"]*)"$`, iShouldNotSee(utils))
	suite.Step(`^(?:|I )should see all of the texts:$`, iShouldSeeAllOfTheTexts(utils))
	suite.Step(`^(?:|I )should see the following(?:| errors| list):$`, iShouldSeeTheFollowing(utils))
	suite.Step(`^(?:|I )should see (?:|the )"([^"]*)"(?: button| link)$`, iShouldSeeButtonLink(utils))
	suite.Step(`^(?:|I )should not see (?:|the )"([^"]*)"(?: button| link)$`, iShouldNotSeeButtonLink(utils))
	suite.Step(`^(?:|I )should see a link that points to "([^"]*)"$`, iShouldSeeALinkThatPointsTo(utils))
	suite.Step(`^(?:|I )should not see a link that points to "([^"]*)"$`, iShouldNotSeeALinkThatPointsTo(utils))
	suite.Step(`^"([^"]*)" should be disabled$`, shouldBeDisabled(utils))
	suite.Step(`^"([^"]*)" should be enabled$`, shouldBeEnabled(utils))
	suite.Step(`^the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)" should be disabled$`, theNthInstanceOfShouldBeDisabled(utils))
	suite.Step(`^the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)" should be enabled$`, theNthInstanceOfShouldBeEnabled(utils))
	suite.Step(`^(?:|I )should see (?:|a |an )"([^"]*)" with "([^"]*)" of "([^"]*)"$`, confirmSeeOfType(utils))
	suite.Step(`^(?:|I )should not see (?:|a |an )"([^"]*)" with "([^"]*)" of "([^"]*)"$`, confirmNotSeeOfType(utils))
	suite.Step(`^(?:|the )"([^"]*)" (?:|element )should exist$`, confirmElementExists(utils))
	suite.Step(`^(?:|the )"([^"]*)" (?:|element )should not exist$`, confirmElementNotExists(utils))
	suite.Step(`^(?:|the )"([^"]*)" (?:|element )should contain "([^"]*)"$`, elementShouldContain(utils))
	suite.Step(`^(?:|the )"([^"]*)" (?:|element )should not contain "([^"]*)"$`, elementShouldNotContain(utils))

	// accessibility steps
	suite.Step(`^(?:|I )test (?:|the page )for accessibility$`, iTestThePageForAccessibility(utils))

	// Check if Scenario Outline has @loadcsv tag
	suite.BeforeScenario(func(s interface{}) {
		if reflect.TypeOf(s) == reflect.TypeOf(&gherkin.ScenarioOutline{}) {
			var err error
			var fileName string
			var csvRecords [][]string
			var csvFileHandle afero.File
			var loadCsvTagExists bool

			// Test for @loadcsv tag
			scenarioOutline := s.(*gherkin.ScenarioOutline)
			flagRegex := regexp.MustCompile(`\@loadcsv\((.+)\)`)

			for _, tag := range scenarioOutline.Tags {
				matches := flagRegex.FindStringSubmatch(tag.Name)
				// use regex to see if the tag matches @loadcsv(...)
				if len(matches) > 1 {
					loadCsvTagExists = true
					fileName = matches[1]
				}
			}

			// If @loadcsv not found, skip
			if !loadCsvTagExists {
				return
			}

			// If no Examples found, skip
			if len(scenarioOutline.Examples) < 1 {
				fmt.Printf("found @loadcsv on Scenario Outline [%s], but no Examples provided...\n", scenarioOutline.Name)
				return
			}

			// Parse CSV file into slice
			if csvFileHandle, err = settings.FileSystem.ResourcesDir.Open(fileName); err != nil {
				fmt.Printf("error encountered while trying to read CSV file [%s]: %s\n", fileName, err)
				return
			}

			csvReader := csv.NewReader(csvFileHandle)
			if csvRecords, err = csvReader.ReadAll(); err != nil {
				fmt.Printf("error encountered parsing CSV file [%s]: %s", fileName, err)
				return
			}

			// Truncate Table Rows before we add the CSV data
			scenarioOutline.Examples[0].TableBody = make([]*gherkin.TableRow, len(csvRecords)-1)

			var csvHeaders = make(map[string]int)
			for idx, row := range csvRecords {
				// Initialize the header map
				if idx == 0 {
					for idx2, header := range row {
						csvHeaders[header] = idx2
					}
				} else {
					tableRow := &gherkin.TableRow{
						Cells: []*gherkin.TableCell{},
					}
					for _, scenarioColumn := range scenarioOutline.Examples[0].TableHeader.Cells {
						csvColumnIdx := csvHeaders[scenarioColumn.Value]
						tableRow.Cells = append(tableRow.Cells, &gherkin.TableCell{
							Value: row[csvColumnIdx],
						})
					}
					scenarioOutline.Examples[0].TableBody[idx-1] = tableRow
				}
			}

			for _, row := range scenarioOutline.Examples[0].TableBody {
				for _, cell := range row.Cells {
					fmt.Printf("%v ", cell.Value)
				}
				fmt.Println()
			}
		}
	})

	// Check for Ajax call in progress
	// An AJAX app should provide a function that returns true if AJAX requests are pending/active:
	// 			window.ajaxPending = function() {
	//
	// 			}
	suite.BeforeStep(func(s *gherkin.Step) {
		var ajaxPending bool
		page.RunScript("return window.ajaxPending() == true;", nil, &ajaxPending)
		for ajaxPending {
			time.Sleep(time.Duration(5) * time.Millisecond)
			page.RunScript("return window.ajaxPending() == true;", nil, &ajaxPending)
		}
	})

	// Get the specified wait times
	waitAfterScenario := utils.Settings.Settings.GetInt("waitAfterScenario")
	waitAfterStep := utils.Settings.Settings.GetInt("waitAfterStep")

	// If a "wait after scenario" is specified, add the pause
	if waitAfterScenario > 0 {
		suite.AfterScenario(func(s interface{}, err error) {
			if err == nil {
				time.Sleep(time.Duration(waitAfterScenario) * time.Millisecond)
			}
		})
	}

	// If a "wait after step" is specified, add the pause
	if waitAfterStep > 0 {
		suite.AfterStep(func(s *gherkin.Step, err error) {
			if err == nil {
				time.Sleep(time.Duration(waitAfterStep) * time.Millisecond)
			}
		})
	}
}

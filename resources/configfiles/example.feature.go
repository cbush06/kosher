package configfiles

// GetExampleFeature returns a sample feature file that can be ran to ensure the new project is configured correctly
func GetExampleFeature() string {
	return `Feature: Test Textbox Manipulation and Validation
    Verify that Kosher's steps for manipulating
    and verifying textboxes work correctly.

    Background:
        Given I maximize the window

    Scenario: Fill a Textbox
        Verify steps for fill a textbox.

        Given I go to the "basicform" page
        When I fill "user-message" with "Hello, World"
        And I press "Show Message"
        Then I should see "Hello, World"

    Scenario: Fill Multiple Textboxes
        Verify steps for filling multiple textboxes.

        Given I go to the "basicform" page
        When I fill in the following:
            | sum1 | 4 |
            | sum2 | 3 |
        And I press "Get Total"
        Then I should see "7"`
}

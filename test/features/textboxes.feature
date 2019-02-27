Feature: Test Textbox Manipulation and Validation
    Verify that Kosher's steps for manipulating
    and verifying textboxes work correctly.

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
        And I pres "Get Total"
        Then I should see "7"

    Scenario: Keying Characters
        Verify that keying characters works.

        Given I go to the "jquery-dropdown" page
        When I click "SelectCountry"
        And I key "u" in "OpenSelectFilter"
        Then I should not see "United States of America"
        And I should see "United States of America"
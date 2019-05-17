Feature: Test Textbox Manipulation and Validation
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

    @docstring
    Scenario: Fill in Textarea
        Verify that filling a textarea with multiline text works.

        Given I go to the "input-form" page
        When I fill "comment" with:
            """
            Hello, World!
            I am Bob.
            """
        Then "comment" should contain:
            """
            Hello, World!
            I am Bob.
            """

    Scenario: Fill Multiple Textboxes
        Verify steps for filling multiple textboxes.

        Given I go to the "basicform" page
        When I fill in the following:
            | sum1 | 4 |
            | sum2 | 3 |
        And I press "Get Total"
        Then I should see "7"

    Scenario: Keying Characters
        Verify that keying characters works.

        Given I go to the "jquery-dropdown" page
        When I click "SelectCountry"
        And I key "u" in "OpenSelectFilter"
        Then I should not see "Bangladesh"
        And I should see "United States of America"
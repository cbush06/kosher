Feature: Test Macros
    Verify that Kosher's macros feature
    works correctly.

    Background:
        Given I maximize the window

    Scenario: Fill a Textbox
        Verify steps for fill a textbox.

        Given I go to the "input-form" page
        When I fill the input form
        Then "first_name" should contain "John"
        And "last_name" should contain "Doe"
        And I fill the rest of the input form
        Then the "Yes" radio should be selected
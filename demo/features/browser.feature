Feature: Test Browser Interaction Features

    Test browser interaction features like screernshots, wait
    maximizing, etc.

    @screenshot
    Scenario: Test Screenshot
        Given I go to the "file-download" page
        And take a screenshot
	And I should see "blahblahblah"

    Scenario: Test Switching Windows
        Given I go to the "popup-window" page
        And I click the "Follow Twitter & Facebook" button
        And I wait 1 second
        When I switch to the 2nd window
        And I maximize the window
        Then I should see "@seleniumeasy"
        And I should see the "Like" button
        When I switch to the 3rd window
        And I maximize the window
        Then I should see "@seleniumeasy"
        And I should see the "Log in and follow @seleniumeasy" button

    Scenario: Test Key in Active Element
        Given I go to the "basicform" page
        And I click "user-message"
        When I key "Hello, World." in the active element
        Then "user-message" should contain "Hello, World."

Feature: Test Browser Interaction Features

    Test browser interaction features like screernshots, wait
    maximizing, etc.

    Scenario: Test Screenshot
        Given I go to the "file-download" page
        And take a screenshot

    Scenario: Test Switching Windows
        Given I go to the "popup-window" page
        And I click the "Follow Twitter & Facebook" button
        When I switch to the 2nd window
        Then I should see "@seleniumeasy"
        And I should see the "Like" button
        When I switch to the 3rd window
        Then I should see "@seleniumeasy"
        And I should see the "Log in and follow @seleniumeasy" button
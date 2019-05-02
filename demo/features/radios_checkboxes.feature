Feature: Test Radio Button and Checkbox Manipulation
    Verify that Kosher's steps for manipulating
    and verifying radio buttons and checkboxes
    work correctly.

    Background:
        Given I maximize the window
        
    Scenario: Selecting a Radio Button
        Given I am on the "radio-button" page
        When I choose the "Female" radio
        And I press "Get Checked value"
        Then I should see "Radio button 'Female' is checked"
        And the "Female" radio should be selected
        And the "Male" radio should not be selected

    Scenario: Check a Single Box
        Given I am on the "basic-checkbox" page
        When I check "Click on this check box"
        Then I should see "Success - Check box is checked"
    
    Scenario: Check Multiple Checkboxes and I Uncheck a Checkbox
        Given I am on the "basic-checkbox" page
        And I wait 2 seconds
        And I should not see the "Uncheck All" button
        And I should see the "Check All" button
        When I check "Option 1"
        And I check "Option 2"
        And I check "Option 3"
        And I check "Option 4"
        And I wait 1 second
        Then I should see the "Uncheck All" button
        And the "Option 4" checkbox should be checked
        And I uncheck "Option 4"
        And I wait 1 second
        Then I should see the "Check All" button
        And the "Option 4" checkbox should not be checked
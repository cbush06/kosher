Feature: Test Radio Button and Checkbox Manipulation
    Verify that Kosher's steps for manipulating
    and verifying radio buttons and checkboxes
    work correctly.

    Scenario: Selecting a Radio Button
        Given I am on the "radio-button" page
        When I choose the "Female" radio
        And I press "Get Checked value"
        Then I should see "Radio button 'Female' is checked"
        And the "Female" radio should be selected
        And the "Male" radio should not be selected
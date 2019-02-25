Feature: Verify Characteristics and Values of a Form

    Verify various characteristics (e.g. disabled fields) and
    values of a form.

    Scenario: Verify Disabled Fields
        Verify that fields are disabled as expected.

        Given I go to the "file-download" page
        Then the first instance of "Generate File" should be disabled
        And "Generate File" should be disabled
    
    Scenario: Verify Enabled Fields
        Verify that fields are enabled as expected.

        Given I go to the "file-download" page
        When I fill "textbox" with "Hello, World."
        Then the first instance of "Generate File" should be enabled
        And "Generate File" should be enabled
    
    Scenario: Verify Today's Date
        Verify a date field contain's todays date.

        Given I go to the "bootstrap-date-picker" page
        When I click "DatePicker"
        And I click the "DatePickerToday" link
        Then "DatePickerTextbox" should contain today's date
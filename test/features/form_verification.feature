Feature: Verify Characteristics and Values of a Form

    Verify various characteristics (e.g. disabled fields) and
    values of a form.

    Background:
        Given I maximize the window

    # Scenario: Verify Disabled Fields
    #     Verify that fields are disabled as expected.

    #     Given I go to the "file-download" page
    #     Then the first instance of "Generate File" should be disabled
    #     And "Generate File" should be disabled

    # Scenario: Verify Enabled Fields
    #     Verify that fields are enabled as expected.

    #     Given I go to the "file-download" page
    #     When I fill "textbox" with "Hello, World."
    #     Then the first instance of "Generate File" should be enabled
    #     And "Generate File" should be enabled

    # Scenario: Verify Selecting Today's Date
    #     Verify a date field contain's todays date after selecting from a date picker.

    #     Given I go to the "bootstrap-date-picker" page
    #     When I click "DatePicker"
    #     And I click the "DatePickerToday" link
    #     Then "DatePickerTextbox" should contain today's date

    # Scenario: Verify Entering Today's Date
    #     Verify a date field contain's todays date after entering it.

    #     Given I go to the "bootstrap-date-picker" page
    #     When I enter today's date in "DatePickerTextbox"
    #     Then "DatePickerTextbox" should contain today's date
    
    # Scenario: Verify Textbox Value
    #     Verify a textbox contains a certain value.

    #     Given I am on the "input-form" page
    #     When I fill "first_name" with "Some"
    #     And I fill "last_name" with "User"
    #     Then "first_name" should contain "Some"
    #     And "first_name" should not contain "User"
    #     And "last_name" should contain "User"
    #     And "last_name" should not contain "Some"
    
    Scenario: Test Send Key
        Verify a new line can be sent using "I send key" step.

        Given I am on the "input-form" page
        When I press the "${ENTER}" key on "comment"
        And I wait 1 second
        And I key "Hello, World" in the "comment" field
        And I press "${BACKSPACE}" key in "comment"
        And I wait 1 second
        And I fill in "comment" with "HELLO WORLD"
        And I wait 2 seconds
        And I press the "${ENTER}" key on "Send"
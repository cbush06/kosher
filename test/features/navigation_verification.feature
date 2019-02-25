Feature: Verify Results of Navigation Actions

    After clicking a link, clicking a button, or submitting
    a form, verfiy that the correct navigations take place.

    Background: Start on the Table Filter page
        Given I go to the "table-search" page

    Scenario: Verify Page
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be on the "bootstrap-date-picker" page

    Scenario: Verify Redirect
        After clicking on a link, verify redirection to the appropriate page.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be redirected to the "bootstrap-date-picker" page
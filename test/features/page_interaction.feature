Feature: Test General Interaction Steps
    Verify that Kosher's general interaction steps
    work as expected.

    Scenario: Indexing Matches
        Verify first element.

        Given I am on the "table-sort" page
        When I press the first instance of "ColumnHeader"
        Then I should see "Y. Berry"
        And I should not see "A. Cox"

    Scenario: Indexing Matches Part 2
        Verify last element.

        Given I am on the "table-sort" page
        When I click the last instance of "ColumnHeader"
        Then I should see "$85,600/y"
        And I should not see "$1,200,000/y"

    Scenario: Accepting JavaScript Alerts
        Verify that alerts can be accepted.

        Given I am on the "js-popup" page
        When I click the first instance of "Click me!"
        And I accept the popup
        Then I should not see the popup "I am an alert box!"

    Scenario: Declining JavaScript Alerts
        Verify that alerts can be declined.

        Given I am on the "js-popup" page
        When I click the first instance of "Click me!"
        And I decline the popup
        Then I should not see the popup "I am an alert box!"

    Scenario: Fill in JavaScript Prompt
        Verfiy that text can be entered in a JavaScript prompt.

        Given I am on the "js-popup" page
        When I click the "Click for Prompt Box" button
        And I enter "Some User" in the popup
        And I accept the popup
        Then I should see "You have entered 'Some User' !"
Feature: Test Select Menu Manipulation
    Verify that Kosher's steps for manipulating
    and verifying select menus works correctly.

    Background:
        Given I maximize the window

    Scenario: Selecting a Single Option from a Dropdown
        Given I am on the "basic-dropdown" page
        When I select "Monday" from "select-demo"
        Then I should see "Day selected :- Monday"

    Scenario: Selecting a Single Option from a Select List
        Given I am on the "basic-dropdown" page
        When I select "New York" from "multi-select"
        And I click "First Selected"
        Then I should see "First selected option is : New York"

    Scenario: Select Multiple Options from a Select List
        Given I am on the "basic-dropdown" page
        When I select the following values from "multi-select":
            | New York     |
            | Pennsylvania |
            | Washington   |
        Then "multi-select" should have the following options selected:
            | New York     |
            | Pennsylvania |
            | Washington   |

    Scenario: Deselect Options from a Select List
        Given I am on the "basic-dropdown" page
        When I select the following values from "multi-select":
            | New York     |
            | Pennsylvania |
            | Washington   |
        And I unselect the following values from "multi-select":
            | Pennsylvania |
        Then "multi-select" should have the following options selected:
            | New York   |
            | Washington |

    Scenario: Verify Available Options
        Given I am on the "basic-dropdown" page
        Then "select-demo" should have the following options:
            | Please select |
            | Sunday        |
            | Monday        |
            | Tuesday       |
            | Wednesday     |
            | Thursday      |
            | Friday        |
            | Saturday      |
        And "multi-select" should have the following options:
            | California   |
            | Florida      |
            | New Jersey   |
            | New York     |
            | Ohio         |
            | Texas        |
            | Pennsylvania |
            | Washington   |
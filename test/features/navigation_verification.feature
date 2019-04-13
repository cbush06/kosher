@Navigation
Feature: Verify Results of Navigation Actions

    After clicking a link, clicking a button, or submitting
    a form, verfiy that the correct navigations take place.

    Background: Start on the Table Filter page
        Given I go to the "table-search" page
        And I maximize the window

    @Smoke
    Scenario: Verify Page
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be on the "bootstrap-date-picker" page

    @Regression
    Scenario: Verify Page with Trailing Slash
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Demo Home" link
        Then I should be on the "home" page

    Scenario: Verify Redirect
        After clicking on a link, verify redirection to the appropriate page.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be redirected to the "bootstrap-date-picker" page

    Scenario: Verify Javascript Alerts
        Verify that alerts can be verified.

        Given I am on the "js-popup" page
        And I should not see the popup "I am an alert box!"
        When I click the first instance of "Click me!"
        And I wait 1 seconds
        Then I should see the popup "I am an alert box!"
        And I accept the popup

    @Smoke @Regression
    Scenario Outline: Verify Redirect
        After clicking on a link, verify redirection to the appropriate page.

        Given I click the "<link1_selector>" link
        And I follow "<link2_selector>" h
        Then I should be redirected to the "<page_name>" page

        @FormsNavigation
        Examples:
        | link1_selector | link2_selector     | page_name      |
        | Input Forms    | Simple Form Demo   | basicform      |
        | Input Forms    | Checkbox Demo      | basic-checkbox |
        | Input Forms    | Radio Buttons Demo | radio-button   |

        @TablesNavigation
        Examples:
        | link1_selector | link2_selector      | page_name    |
        | Table          | Table Data Search   | table-search |
        | Table          | Table Sort & Search | table-sort   |
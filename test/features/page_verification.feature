Feature: Verify Miscellaneous Elements of the Page

    Verify that Kosher's steps for verifying
    various elements on the page work properly.

    Scenario: Confirm I see all the Texts
        Verify that I see various labels and
        excerpts of text on the page.

        Given I go to the "jquery-dropdown" page
        Then I should see all of the texts:
            | Select Country                 |
            | Select State                   |
            | Select US Outlying Territories |

    Scenario: Confirm Buttons/Links are Visible on the Page
        Verify that I see a button and/or link on the page
        and that I do not see a button and/or link on the page.

        Given I go to the "table-search" page
        When I click "Date pickers"
        Then I should see the "Bootstrap Date Picker" link
        And I should not see the "Table Pagination" link
        And I should see the "Filter" button
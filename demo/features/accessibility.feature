Feature: Test Kosher's Accessibility Scanning

   This feature tests Kosher's use of the Axe-Core API
   to scan pages for Accessibility issues.

   Scenario: Scan a page.
        Given I am on the "bar-chart" page
        And I test the page for accessibility
    
    Scenario: Scan another page
        Given I am on the "pie-chart" page
        And I test the page for accessibility
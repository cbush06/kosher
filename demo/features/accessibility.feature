Feature: Test Kosher's Accessibility Scanning

   This feature tests Kosher's use of the Axe-Core API
   to scan pages for Accessibility issues.

   Scenario: Scan a page.
        Given I am on the "home" page
        And I test the page for accessibility
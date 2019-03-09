---
layout: default
title: progress Report Format
description: Description of the "progress" report format.
parent: Reports
nav_order: 6
---

# progress Report Format

The `progres` report format prints the real-time progress of test execution to the console.
{: .fs-6 .fw-300 }

## Foramt

The `progress` report format marks each step executed with a period in the terminal followed by summary statistics for the test execution.

For the `*.feature` file:

```gherkin
Feature: Verify Results of Navigation Actions

    After clicking a link, clicking a button, or submitting
    a form, verfiy that the correct navigations take place.

    Background: Start on the Table Filter page
        Given I go to the "table-search" page
        And I maximize the window

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

    Scenario: Verify Javascript Alerts
        Verify that alerts can be verified.

        Given I am on the "js-popup" page
        And I should not see the popup "I am an alert box!"
        When I click the first instance of "Click me!"
        And I wait 1 seconds
        Then I should see the popup "I am an alert box!"
```

The `progress` format would generate output similar to this:

![progress report format]({{site.baseurl}}/assets/images/progress.png)
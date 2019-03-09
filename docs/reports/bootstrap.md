---
layout: default
title: bootstrap Report Format
description: Description of the "bootstrap" report format.
parent: Reports
nav_order: 1
---

# bootstrap Report Format

The `bootstrap` report format produces an HTML file complete with charts, summary statistics, and granular reporting on feature files, scenarios, and steps executed during the test. It is the _most complete_ report available.
{: .fs-6 .fw-300 }

## Format

The `bootstrap` report format generates a beautiful HTML report. The format is tailored to export nicely to PDF, but **does not** import well into Word.

The report displays the `project name` in the upper left-hand corner, summary statistics in the upper right-hand corner, and the date of test execution below those statistics.

Next, it provides 2 charts: the left chart shows failed vs. passed scenarios; the right chart shows passed, failed, skipped, and pending scenarios. _Pending_ scenarios are those that could not be matched to a known kosher step (likely, these represent typos).

After the charts, meta data about the project is shown.

The final section of the report shows the statuses of feature files, scenarios, and steps.

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

The `bootstrap` report format would generate a `results.html` file similar to this:

![bootstrap report format]({{site.baseurl}}/assets/images/bootstrap.png)
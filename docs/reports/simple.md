---
layout: default
title: simple Report Format
description: Description of the "simple" report format.
parent: Reports
nav_order: 6
---

# simple Report Format

The `simple` report format produces an HTML file that is easily imported by Word and exported to PDF.
{: .fs-6 .fw-300 }

## Format

The `simple` format generates an HTML file that displays summary statistics followed by the results of each `*.feature` file, scenario, and step executed.

The report displays the `project name` in the upper left-hand corner, summary statistics in the upper right-hand corner, and the date of test execution below those statistics.

Next, it provides meta data about the project is shown.

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

The `simple` report format would generate a `results.html` file similar to this:

![simple report format]({{site.baseurl}}/assets/images/simple.png)
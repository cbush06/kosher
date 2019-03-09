---
layout: default
title: pretty Report Format
description: Description of the "pretty" report format.
parent: Reports
nav_order: 4
---

# pretty Report Format

The `pretty` report format prints real-time results of features, scenarios, and steps to the terminal.
{: .fs-6 .fw-300 }

## Format

The `pretty` format generates color-coded test execution results and prints them to the console in real time. It provides summarized statistics at the end of the test.

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

The `pretty` report format would print this to the terminal:

![pretty report format]({{site.baseurl}}/assets/images/pretty.png)
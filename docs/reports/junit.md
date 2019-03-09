---
layout: default
title: junit Report Format
description: Description of the "junit" report format.
parent: Reports
nav_order: 4
---

# junit Report Format

The `junit` report format produces an XML file in the JUnit report format that records summary statistics for the feature files and scenarios executed.
{: .fs-6 .fw-300 }

# Format

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

The `junit` format would produce a `results.xml` file similar to this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<testsuites name="kosher" tests="3" skipped="0" failures="0" errors="0" time="6.765098222s">
  <testsuite name="Verify Results of Navigation Actions" tests="3" skipped="0" failures="0" errors="0" time="6.763860018s">
    <testcase name="Verify Page" status="passed" time="3.527778577s"></testcase>
    <testcase name="Verify Redirect" status="passed" time="850.200106ms"></testcase>
    <testcase name="Verify Javascript Alerts" status="passed" time="2.385874047s"></testcase>
  </testsuite>
</testsuites>
```
---
layout: default
title: bootstrap Report Format
description: Description of the "bootstrap" report format.
parent: Reports
nav_order: 1
---

# Accessibility Report

The `accessibility` report is an HTML file complete with charts, summary statistics, and granular reporting on on all accessibility findings produced by the [`I test the page for accessibility`](../steps/i_test_for_accessibility.html) step.
{: .fs-6 .fw-300 }

## Format

The `accessibility` report generates a beautiful HTML report. The format is tailored to export nicely to PDF, but **does not** import well into Word.

The report displays the `project name` in the upper left-hand corner, total issues by impact level in the upper right-hand corner, and the date of test execution below those statistics.

Next, it provides 2 charts: the left chart shows total violation issues vs. issues needing review; the right chart shows the number of issues per each violation level found. Issues *needing review* are issues that the Axe accessibility scanner could not determine to definitively pass or fail a certain test.

After the charts, meta data about the project and accessibility scan/scanner shown.
* **App Version:** the `appVersion` you set in [settings.json]([../config/settings.html](../config/settings.html#appversion))
* **OS:** the operating system the test was ran on
* **Environment:** either the `defaultEnvironment` you set in [settings.json](../config/settings.html#defaultenvironment) or the command-line arguments [`-e` or `--environment`](https://cbush06.github.io/kosher/cli.html#arguments-1)
* **Browser:** the browser driver selected by `driver` in [settings.json](../config/settings.html#driver)
* **Axe Version:** the version of the [axe-core API](https://github.com/dequelabs/axe-core) kosher is using
* **Rule Sets:** [Axe tags](https://www.deque.com/axe/axe-for-web/documentation/api-documentation/#api-name-axegetrules) specified by [`ruleSets`](../config/settings.html#rulesets)
* **Threshold:** the minimum severity level any finding must be before kosher fails the `I test the page for accessibility` step; this is defined by [`impactThreshold`](../config/settings.html#impactthreshold)

The final section of the report groups findings into the following hiearchy (from most general to most specific):
* pages scanned
* rules enforced on pages
* elements found violating rules

The following information is provided for each element found to violate a rule:
* CSS query that will find that specific element
* HTML excerpt of that element
* Suggested fixes to bring the element into compliance with the rule

For the `*.feature` file:

```gherkin
Feature: Test Kosher's Accessibility Scanning

   This feature tests Kosher's use of the Axe-Core API
   to scan pages for Accessibility issues.

   Scenario: Scan a page.
        Given I am on the "bar-chart" page
        And I test the page for accessibility
    
    Scenario: Scan another page
        Given I am on the "pie-chart" page
        And I test the page for accessibility
```

The `accessibility` report would look similar to this:

![accessibility report]({{site.baseurl}}/assets/images/accessibility_report.png)
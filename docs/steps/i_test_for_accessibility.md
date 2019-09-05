---
layout: default
title: I test the page for accessibility
parent: Step Definitions
---

# I test the page for accessibility

Test the page for accessibility compliance according to the rule sets defined in [`settings.json`](../config/settings.html). This step will fail if any accessibility findings have an impact (i.e. severity) rating that equals or exceeds the [`impactThreshold`](../config/settings.html#impactThreshold) setting.
{: .fs-6 .fw-300 }

If this step is ran at least once, kosher will produce 2 accessibility results files:
* [**axe-results.html:**](../reports/accessibility.html) a bootstrap formatted report of all accessibility findings including two charts that summarize the number and severity of issues found
* **axe-results.json:** a JSON file that stores findings in a machine-friendly format kosher uses when you send test results to other systems (e.g. JIRA)

## Pattern

```
(?:|I )test (?:|the page )for accessibility
```

## Parameters

None.

## Examples

```gherkin
Given I am on the "home" page
And I test the page for accessibility
And go to the "contact" page
And I test for accessibility
```
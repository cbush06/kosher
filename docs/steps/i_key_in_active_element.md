---
layout: default
title: "I key \"__\" in the active element"
parent: Step Definitions
---

# I key "\_\_" in the active element

Simulate a user keying a value into/onto the currently active element within the browser.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )key "([^"]*)"(?:| in the active element)$
```

## Parameters

| Position | Description | Value Type                            | Restrictions                                                                             |
| :------: | ----------- | ------------------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | value       | string/[variables](../variables.html) |                                                                                          |

## Examples

```gherkin
Given I switch to iframe 1
And I key "Hello, World" in the active element
And I key "Goodbye, World"
```
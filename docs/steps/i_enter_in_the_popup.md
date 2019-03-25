---
layout: default
title: "I enter \"__\" in the popup"
parent: Step Definitions
---

# I enter "\_\_" in the popup

Fill in a form field with a value.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )enter "([^"]*)" in the popup$
```

## Parameters

| Position | Description | Value Type                            | Restrictions |
| :------: | ----------- | ------------------------------------- | ------------ |
|    1     | value       | string/[variables](../variables.html) |              |

## Examples

```gherkin
When I click the "Show Popup" button
And I enter "my name" in the popup
```
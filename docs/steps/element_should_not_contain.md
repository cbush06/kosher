---
layout: default
title: "the \"__\" element should not contain \"__\""
parent: Step Definitions
---

# the "\_\_" element should not contain "\_\_"

Verifies that an element matched by the selector, ID, name, or label does not contain the specified text.
{: .fs-6 .fw-300 }

## Pattern

```
^(?|the )"([^"]*)" (?|element )should not contain "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                            | Restrictions |
| :------: | ----------- | ------------------------------------- | ------------ |
|    1     | field       | field id/name/label/selector          |              |
|    2     | value       | string/[variables](../variables.html) |              |

## Examples

```gherkin
Given I am on the "home" page
Then the "HomeCenterMessage" element should not contain "Hello, World."
```
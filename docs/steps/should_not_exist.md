---
layout: default
title: "the \"__\" element should not exist"
parent: Step Definitions
---

# the "\_\_" element should exist

Verifies that the page does not contain an element matched by the selector, ID, name, or label.
{: .fs-6 .fw-300 }

## Pattern

```
^(?|the )"([^"]*)" (?|element )should not exist$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | field       | field id/name/label/selector |              |

## Examples

```gherkin
Given I am on the "register" page
Then the "search" element should not exist
```
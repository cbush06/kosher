---
layout: default
title: "I hover over \"__\""
parent: Step Definitions
---

# I hover over "\_\_"

Simulate a mouseover action on the specified element.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )hover over "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | element     | field id/name/label/selector |              |

## Examples

```gherkin
When I hover over "Departments"
And I click the "Men's Wear" link
```
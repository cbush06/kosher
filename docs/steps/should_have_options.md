---
layout: default
title: "\"__\" should have the following options:"
parent: Step Definitions
---

# "\_\_" should have the following options:

Verifies that a select list has **exactly** the specified options.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should have (?:|the )following options:$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                     |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------- |
|    1     | select      | field id/name/label/selector | Must refer to a [select field]({{site.baseurl}}/field_types.html#select-fields). |

## Data Table Columns

| Column Position | Description | Value Type | Restrictions |
| :-------------: | ----------- | ---------- | ------------ |
|        1        | value       | string     |              |

## Examples

```gherkin
Then "Pizza Toppings" should have the following options:
    | pepperonis   |
    | bell peppers |
    | onions       |
    | hamburger    |
    | black olives |
And "Condiments" should have following options:
    | Italian seasoning           |
    | Garlic butter dipping sauce |
    | Ranch dressing              |
    | Marinara sauce              |
```
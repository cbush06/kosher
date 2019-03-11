---
layout: default
title: "I unselect the following values from \"__\":"
parent: Step Definitions
---

# I unselect the following values from "\_\_":

Unselects one or more values from a multi-select list.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )unselect (?:|the )following values from "([^"]*)":$
```

## Parameters

| Position | Description       | Value Type                   | Restrictions                                                                                                       |
| :------: | ----------------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------ |
|    1     | multi-select list | field id/name/label/selector | Must refer to a [select field]({{site.baseurl}}/field_types.html#select-fields) with the `multiple` attribute set. |

## Data Table Columns

| Column Position | Description | Value Type | Restrictions |
| :-------------: | ----------- | ---------- | ------------ |
|        1        | value       | string     |              |

## Examples

```gherkin
When I unselect the following values from "Pizza Toppings":
    | pepperonis   |
    | bell peppers |
    | onions       |
    | hamburger    |
    | black olives |
And unselect following values from "Condiments":
    | Italian seasoning           |
    | Garlic butter dipping sauce |
    | Ranch dressing              |
    | Marinara sauce              |
```
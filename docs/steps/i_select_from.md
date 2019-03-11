---
layout: default
title: I select "__" from "__"
parent: Step Definitions
---

# I select "\_\_" from "\_\_"

Select a single value from a single- or multi-select list.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )select "([^"]*)" from "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                           |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [select field]({{site.baseurl}}/field_types.html#select-fields).       |
|    2     | value       | string                       | Must match the actual `value` attribute of the `<option></option>` you wish to select. |

## Examples

```gherkin
When I select "Alabama" from "state"
And select "Birmingham" from "city"
```
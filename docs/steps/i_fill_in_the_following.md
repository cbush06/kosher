---
layout: default
title: "I fill in the following:"
parent: Step Definitions
---

# I fill in the following:

Fill in multiple fields with a single step.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )fill in the following:$^(?:|I )check "([^"]*)"$
```

## Parameters

None.

## Data Table Columns

| Column Position | Description        | Value Type                   | Restrictions                                                                                                                                                                                                                                                                                  |
| :-------------: | ------------------ | ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|        1        | field to be filled | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields), [radio button]({{site.baseurl}}/field_types.html#selection-fields), [checkbox]({{site.baseurl}}/field_types.html#selection-fields), or [select]({{site.baseurl}}/field_types.html#selection-fields). |
|        2        | value              | string/number/boolean        | If the field is a radio button, value must be either `true` or `false` (meaning _checked_ or _unchecked_, respectively).                                                                                                                                                                      |

## Examples

```gherkin
When I fill in the following:
| First Name                 | John                                    |
| zip_code                   | 29339                                   |
| cb_receive_mailings        | true                                    |
| age                        | 34                                      |
| How did you hear about us? | My friend Joe gave me your web address. |
```
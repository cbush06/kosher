---
layout: default
title: "\"__\" should not contain:"
parent: Step Definitions
---

# "\_\_" should not contain:

Verifies that the specified field (presumably a textarea) does not contain the provided multiline value.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should not contain:$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Docstring

The field's value will be compared to the docstring following the step.

## Examples

```gherkin
Then "Remarks" should not contain:
"""
Four score and seven
years ago our fathers
brought forth, on this
continent, a new nation
```
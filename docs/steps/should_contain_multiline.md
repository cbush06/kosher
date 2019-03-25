---
layout: default
title: "\"__\" should contain:"
parent: Step Definitions
---

# "\_\_" should contain:

Verifies that the specified field (presumably a textarea) contains the provided multiline value.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should contain:$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Docstring

The field's value will be compared to the docstring following the step. The docstring may include [variables](../variables.html).

## Examples

```gherkin
Then "Remarks" should contain:
"""
Four score and seven
years ago our fathers
brought forth, on this
continent, a new nation
```
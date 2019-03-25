---
layout: default
title: "I fill in \"__\" with:"
parent: Step Definitions
---

# I fill in "\_\_" with:

Fill in a form field (presumably a textarea) with a multiline value.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )fill (?:|in )"([^"]*)" with:"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                                                                                                                                                                                                                                                               |
| :------: | ----------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields), [radio button]({{site.baseurl}}/field_types.html#radio-button-fields), [checkbox]({{site.baseurl}}/field_types.html#checkbox-fields), or [select]({{site.baseurl}}/field_types.html#select-fields).                                               |

## Docstring

The docstring following the step will be used to fill the field. The docstring may include [variables](../variables.html).

## Examples

```gherkin
When I fill in "Comments" with:
"""
kosher is an amazing tool.
It has saved us countless
ours doing regression testing
and serves as a single source of
truth for requirements.
"""
```
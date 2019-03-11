---
layout: default
title: I fill in "__" with "__"
parent: Step Definitions
---

# I fill in "\_\_" with "\_\_"

Fill in a form field with a value.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )fill (?:|in )"([^"]*)" with "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                                                                                                                                                                                                                 |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields), [radio button]({{site.baseurl}}/field_types.html#radio-button-fields), [checkbox]({{site.baseurl}}/field_types.html#checkbox-fields), or [select]({{site.baseurl}}/field_types.html#select-fields). |
|    2     | value       | string/boolean               | If the field is a radio button, value must be either `true` or `false` (meaning _checked_ or _unchecked_, respectively).                                                                                                                                                                     |

## Examples

```gherkin
When I fill in "City" with "Los Angeles"
And fill "Accept Terms" with "true"
```
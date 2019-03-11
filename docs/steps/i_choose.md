---
layout: default
title: I choose "__"
parent: Step Definitions
---

# I choose "\_\_"

Set a checkbox to the <em>checked</em> state.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )choose "([^"]*)"$
```

## Parameters

| Position | Description  | Value Type                   | Restrictions                                                                                 |
| :------: | ------------ | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | radio button | field id/name/label/selector | Must refer to a [radio button field]({{site.baseurl}}/field_types.html#radio-button-fields). |

## Examples

```gherkin
When I choose "male"
And I choose "18-25"
```
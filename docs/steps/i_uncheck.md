---
layout: default
title: I uncheck "__"
parent: Step Definitions
---

# I uncheck "\_\_"

Set a checkbox to the <em>unchecked</em> state.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )uncheck "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                         |
| :------: | ----------- | ---------------------------- | ------------------------------------------------------------------------------------ |
|    1     | checkbox    | field id/name/label/selector | Must refer to a [checkbox field]({{site.baseurl}}/field_types.html#checkbox-fields). |

## Examples

```gherkin
When I uncheck "Accept terms"
And uncheck "Receive newsletters"
```
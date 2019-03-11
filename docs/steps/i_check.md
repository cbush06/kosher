---
layout: default
title: I check "__"
parent: Step Definitions
---

# I check "\_\_"

Set a checkbox to the <em>checked</em> state.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )check "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                         |
| :------: | ----------- | ---------------------------- | ------------------------------------------------------------------------------------ |
|    1     | checkbox    | field id/name/label/selector | Must refer to a [checkbox field]({{site.baseurl}}/field_types.html#checkbox-fields). |

## Examples

```gherkin
When I check "Accept terms"
And check "Receive newsletters"
```
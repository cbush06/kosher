---
layout: default
title: "the \"__\" checkbox should be checked"
parent: Step Definitions
---

# the "\_\_" checkbox should be checked

Verifies that the specified checkbox is checked.
{: .fs-6 .fw-300 }

## Pattern

```
^the "([^"]*)" checkbox should be checked$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                         |
| :------: | ----------- | ---------------------------- | ------------------------------------------------------------------------------------ |
|    1     | field       | field id/name/label/selector | Must refer to a [checkbox field]({{site.baseurl}}/field_types.html#checkbox-fields). |

## Examples

```gherkin
Then the "Accept terms" checkbox should be checked
```
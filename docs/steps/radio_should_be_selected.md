---
layout: default
title: "the \"__\" radio should be selected"
parent: Step Definitions
---

# the "\_\_" radio should be selected

Verifies that the specified radio button is selected.
{: .fs-6 .fw-300 }

## Pattern

```
^the "([^"]*)" radio should be selected$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                 |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [radio button field]({{site.baseurl}}/field_types.html#radio-button-fields). |

## Examples

```gherkin
Then the "Model 3 Tesla" radio should be selected
```
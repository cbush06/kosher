---
layout: default
title: "the \"__\" radio should not be selected"
parent: Step Definitions
---

# the "\_\_" radio should not be selected

Verifies that the specified radio button is unselected.
{: .fs-6 .fw-300 }

## Pattern

```
^the "([^"]*)" radio should not be selected$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                 |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [radio button field]({{site.baseurl}}/field_types.html#radio-button-fields). |

## Examples

```gherkin
Then the "Chevy Volt" radio should not be selected
```
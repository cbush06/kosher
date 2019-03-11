---
layout: default
title: "\"__\" should not contain \"__\""
parent: Step Definitions
---

# "\_\_" should not contain "\_\_"

Verifies that the specified field does not contain the provided value.
{: .fs-6 .fw-300 }

## Pattern

```golang
^"([^"]*)" should not contain "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |
|    2     | value       | string                       |                                                                                          |

## Examples

```gherkin
Then "Approval Status" should not contain "Denied"
```
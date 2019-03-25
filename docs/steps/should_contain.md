---
layout: default
title: "\"__\" should contain \"__\""
parent: Step Definitions
---

# "\_\_" should contain "\_\_"

Verifies that the specified field contains the provided value.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should contain "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                            | Restrictions                                                                             |
| :------: | ----------- | ------------------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector          | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |
|    2     | value       | string/[variables](../variables.html) |                                                                                          |

## Examples

```gherkin
Then "Approval Status" should contain "Approved"
```
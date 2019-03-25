---
layout: default
title: "I key in \"__\" in the \"__\" field"
parent: Step Definitions
---

# I key in "\_\_" in the "\_\_" field

Simulate a user keying a value into a field.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )key (?:|in )"([^"]*)" in (?:|the )"([^"]*)"(?:| field)$
```

## Parameters

| Position | Description | Value Type                            | Restrictions                                                                             |
| :------: | ----------- | ------------------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | value       | string/[variables](../variables.html) |                                                                                          |
|    2     | field       | field id/name/label/selector          | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Examples

```gherkin
When I key in "John" in the "First Name" field
And I key "Doe" in the "Last Name" field
And key "green" in "Favorite Color"
```
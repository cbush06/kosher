---
layout: default
title: "\"__\" should contain today's date"
parent: Step Definitions
---

# "\_\_" should contain today's date

Parses the value the field's value as a date using the `dateFormat` in the `settings.json` file and verifies that it is equal to the current date.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should contain today's date$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Examples

```gherkin
Then "Contract Effective Date" should contain today's date
```
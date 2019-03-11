---
layout: default
title: I verify "__" has today's date
parent: Step Definitions
---

# I verify "\_\_" has today's date

Parses the value the field's value as a date using the `dateFormat` in the `settings.json` file and verifies that it is equal to the current date.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )verify "([^"]*)" has today's date$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Examples

```gherkin
When I "Contract Effective Date" has today's date
And verify "Signature Date" has today's date
```
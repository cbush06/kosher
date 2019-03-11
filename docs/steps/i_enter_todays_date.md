---
layout: default
title: I enter today's date in "__"
parent: Step Definitions
---

# I enter today's date in "\_\_"

Enter the current date in the specified field using the `dateFormat` format in the `settings.json` file.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )enter today's date in "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                             |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------------- |
|    1     | field       | field id/name/label/selector | Must refer to a [text-based field]({{site.baseurl}}/field_types.html#text-based-fields). |

## Examples

```gherkin
When I enter today's date in "Contract Effective Date"
And enter today's date in "Signature Date"
```
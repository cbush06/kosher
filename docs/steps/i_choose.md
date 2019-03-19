---
layout: default
title: "I choose the \"__\" radio"
parent: Step Definitions
---

# I choose the "\_\_" radio

Set a radio button to the <em>checked</em> state.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )choose (?:|the )"([^"]*)" radio$
```

## Parameters

| Position | Description  | Value Type                   | Restrictions                                                                                 |
| :------: | ------------ | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | radio button | field id/name/label/selector | Must refer to a [radio button field]({{site.baseurl}}/field_types.html#radio-button-fields). |

## Examples

```gherkin
When I choose the "male" radio
And choose "18-25" radio
```
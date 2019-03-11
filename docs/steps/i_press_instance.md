---
layout: default
title: "I press/click the __ instance of \"__\""
parent: Step Definitions
---

# I press/click the \_\_ instance of "\_\_"

Given multiple elements that match parameter 2, press/click the n<sup>th</sup> element as specified by parameter 1.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )(?:press|click) the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)"$
```

## Parameters

| Position | Description           | Value Type                   | Restrictions                                           |
| :------: | --------------------- | ---------------------------- | ------------------------------------------------------ |
|    1     | positional identifier | string/numeric               | May be `first`, `last`, `#th`, `#st`, `#rd`, or `#nd`. |
|    2     | element               | field id/name/label/selector |                                                        |

## Examples

```gherkin
Given I press the first instance of "paginatorLinks"
And click the 2nd instance of "paginatorLinks"
And click the 3rd instance of "paginatorLinks"
When I press the 4th instance of "rowEditButtons"
And click the 21st instance of "editChildRecordLinks"
```
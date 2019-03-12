---
layout: default
title: "the __ instance of \"__\" should be enabled"
parent: Step Definitions
---

# the \_\_ instance of "\_\_" should be enabled

Verifies that the n<sup>th</sup> match of the specified selector is enabled.
{: .fs-6 .fw-300 }

## Pattern

```
^the (first|last|[0-9]+(?:th|st|rd|nd)) instance of "([^"]*)" should be enabled$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | field       | field id/name/label/selector |              |

## Examples

```gherkin
Given I am on the "page3Results" page
Then the 3rd instance of "paginationButton" should be disabled
And the 2nd instance of "paginationButton" should be enabled
```
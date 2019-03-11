---
layout: default
title: "\"__\" should be disabled"
parent: Step Definitions
---

# "\_\_" should be disabled

Verifies that the specified field is disabled.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should be disabled$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | field       | field id/name/label/selector |              |

## Examples

```gherkin
Then "Continue" should be disabled
```
---
layout: default
title: "\"__\" should be enabled"
parent: Step Definitions
---

# "\_\_" should be enabled

Verifies that the specified field is enabled.
{: .fs-6 .fw-300 }

## Pattern

```
^"([^"]*)" should be enabled$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | field       | field id/name/label/selector |              |

## Examples

```gherkin
Then "Continue" should be enabled
```
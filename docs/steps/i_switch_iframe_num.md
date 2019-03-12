---
layout: default
title: I switch to iframe __
parent: Step Definitions
---

# I switch to iframe \_\_

Switch control to the n<sup>th</sup> `<iframe></iframe>` element.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )switch to iframe (\d+)$
```

## Parameters

| Position | Description           | Value Type | Restrictions                |
| :------: | --------------------- | ---------- | --------------------------- |
|    1     | positional identifier | integer    | Must be a positive integer. |

## Examples

```gherkin
Given I switch to iframe 2
```
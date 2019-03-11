---
layout: default
title: I switch to frame __
parent: Step Definitions
---

# I switch to frame \_\_

Switch control to the n<sup>th</sup> `<frame></frame>` element.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )switch to frame (\d+)$
```

## Parameters

| Position | Description           | Value Type | Restrictions                |
| :------: | --------------------- | ---------- | --------------------------- |
|    1     | positional identifier | integer    | Must be a positive integer. |

## Examples

```gherkin
Given I switch to frame 2
```
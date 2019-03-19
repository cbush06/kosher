---
layout: default
title: I switch to the __ window
parent: Step Definitions
---

# I switch to the \_\_ window

Switch control to the n<sup>th</sup> open window/tab.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )switch to the (first|last|[0-9]+(?:th|st|rd|nd)) window$	
```

## Parameters

| Position | Description           | Value Type | Restrictions                                           |
| :------: | --------------------- | ---------- | ------------------------------------------------------ |
|    1     | positional identifier | string     | May be `first`, `last`, `#th`, `#st`, `#rd`, or `#nd`. |

## Examples

```gherkin
Given I switch to first window
And switch to the 2nd window
And switch to the 3rd window
And switch to the 4th window
And switch to the 21st window
And switch to the last window
```
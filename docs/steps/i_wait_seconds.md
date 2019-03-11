---
layout: default
title: I wait __ seconds
parent: Step Definitions
---

# I wait \_\_ seconds

Force test execution to pause for a specified number of seconds. This is often useful when you must wait for an AJAX request/response or long-running calculation to complete. **Note** that you do not need to "wait" for normal browser navigations, as test execution is automatically paused until these complete.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )wait (\d+) second(?:|s)$
```

## Parameters

| Position | Description | Value Type | Restrictions                |
| :------: | ----------- | ---------- | --------------------------- |
|    1     | second(s)   | integer    | Must be a positive integer. |

## Examples

```gherkin
When I wait 1 second
And press the "Refresh Data" button
And wait 5 seconds
```
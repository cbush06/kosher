---
layout: default
title: I take a screenshot
parent: Step Definitions
---

# I take a screenshot

Take a screenshot of the browser's rendered web page. The screenshot will be named with the following pattern `DDMMMYYYY-HHMMSS.MIL.png` (where `MIL` is a 3 digit millisecond value). Screenshots are saved in the `/results` directory of a project.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )take a screenshot$
```

## Parameters

None.

## Examples

```gherkin
When I take a screenshot
And go to the "contact" page
And take a screenshot
```
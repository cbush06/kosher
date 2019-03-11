---
layout: default
title: I maximize the window
parent: Step Definitions
---

# I maximize the window

Resize the window to the dimensions set in the `settings.json` file for the currently selected [`screenFormat`]({{site.baseurl}}/config/settings.html#screenformat). **Note** that this step is usually unnecessary because kosher automatically resizes the window upon startup.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )maximize the window$
```

## Parameters

None.

## Examples

```gherkin
Given I maximize the window
```
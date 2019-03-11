---
layout: default
title: I accept the popup
parent: Step Definitions
---

# I accept the popup

Affirmatively responds to the presently shown [browser pop-up]({{site.baseurl}}/field_types.html#browser-pop-ups). **Note** that this step will fail if there is not a pop-up shown in the browser.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )accept the popup$
```

## Parameters

None.

## Examples

```gherkin
When I accept the popup
```
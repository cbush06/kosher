---
layout: default
title: I decline the popup
parent: Step Definitions
---

# I decline the popup

Negatively responds to the presently shown [browser pop-up]({{site.baseurl}}/field_types.html#browser-pop-ups). **Note** that this step will fail if there is not a pop-up shown in the browser.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )decline the popup$
```

## Parameters

None.

## Examples

```gherkin
When I decline the popup
```
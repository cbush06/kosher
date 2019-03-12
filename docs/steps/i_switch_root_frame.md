---
layout: default
title: I switch to the root frame __
parent: Step Definitions
---

# I switch to frame \_\_

Switch control to the original, root frame that was activated when the page originally loaded (i.e. before any `I switch to iframe __` or `I switch to frame __` steps were executed).
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )switch to the root frame$
```

## Parameters

None.

## Examples

```gherkin
Given I switch to frame 2
And I click the "Contact Us" link
When I switch to the root frame
Then I should be on the "contactUs" page
```
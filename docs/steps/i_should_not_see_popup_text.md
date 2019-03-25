---
layout: default
title: "I should not see the popup text \"__\""
parent: Step Definitions
---

# I should not see the popup text "\_\_"

Verifies that the presently shown [browser pop-up]({{site.baseurl}}/field_types.html#browser-pop-ups) does not contain the specified text. **Note** that this step will fail if there is not a pop-up shown in the browser.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should not see (?:|the )popup (?:|text )"([^"]*)"
```

## Parameters

| Position | Description                         | Value Type | Restrictions |
| :------: | ----------------------------------- | ---------- | ------------ |
|    1     | text/[variables](../variables.html) | string     |              |

## Examples

```gherkin
Given I am on the "google-home" page
When I click the "About" link
Then I should not see the popup text "Yahoo, Inc. Version 1,223,961,250"
```
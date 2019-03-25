---
layout: default
title: "I should see the popup text \"__\""
parent: Step Definitions
---

# I should see the popup text "\_\_"

Verifies that the presently shown [browser pop-up]({{site.baseurl}}/field_types.html#browser-pop-ups) contains the specified text. **Note** that this step will fail if there is not a pop-up shown in the browser.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should see (?:|the )popup (?:|text )"([^"]*)"
```

## Parameters

| Position | Description | Value Type                            | Restrictions |
| :------: | ----------- | ------------------------------------- | ------------ |
|    1     | text        | string/[variables](../variables.html) |              |

## Examples

```gherkin
Given I am on the "google-home" page
When I click the "About" link
Then I should see the popup text "Google, Inc. Version 1,223,961,250"
```
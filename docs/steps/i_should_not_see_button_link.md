---
layout: default
title: "I should not see the \"__\" button/link"
parent: Step Definitions
---

# I should not see the "\_\_" button/link

Verifies that the page does not contain a button or link with the specified label.
{: .fs-6 .fw-300 }

## Pattern

```
^(?: |I )should not see (?:|the )"([^"]*)"(?: button| link)$
```

## Parameters

| Position | Description | Value Type | Restrictions |
| :------: | ----------- | ---------- | ------------ |
|    1     | text        | string     |              |

## Examples

```gherkin
Given I am on the "home" page
Then I should not see the "Log Out" link
```
---
layout: default
title: "I should see the \"__\" button/link"
parent: Step Definitions
---

# I should see the "\_\_" button/link

Verifies that the page contains a button or link with the specified label.
{: .fs-6 .fw-300 }

## Pattern

```
^(?: |I )should see (?:|the )"([^"]*)"(?: button| link)$
```

## Parameters

| Position | Description | Value Type | Restrictions |
| :------: | ----------- | ---------- | ------------ |
|    1     | text        | string     |              |

## Examples

```gherkin
Given I am on the "home" page
Then I should see the "Log Out" link
```
---
layout: default
title: "I should not see \"__\""
parent: Step Definitions
---

# I should not see "\_\_"

Verifies that the page does not contain the specified text. **Note** that this step may sometimes fail if the text is split across multiple elements (e.g. parts of the text are styled individually using multiple `<span></span>` elements).
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should not see "([^"]*)"$
```

## Parameters

| Position | Description                         | Value Type | Restrictions |
| :------: | ----------------------------------- | ---------- | ------------ |
|    1     | text/[variables](../variables.html) | string     |              |

## Examples

```gherkin
Given I am on the "contactUs" page
Then I should not see "Catalog Contents"
```
---
layout: default
title: "I should see all of the texts:"
parent: Step Definitions
---

# I should see all of the texts:

Verifies that the page contains all of the specified texts. **Note** that this step may sometimes fail if any of the texts are split across multiple elements (e.g. parts of a text are styled individually using multiple `<span></span>` elements).
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should see all of the texts:$
```

## Parameters

None.

## Data Table Columns

| Column Position | Description                         | Value Type | Restrictions |
| :-------------: | ----------------------------------- | ---------- | ------------ |
|        1        | text/[variables](../variables.html) | string     |              |

## Examples

```gherkin
Given I am on the "home" page
Then I should see all of the texts:
    | Catalog         |
    | Specials        |
    | Discounts       |
    | Become a Member |
    | Contact Us      |
    | About Us        |
```
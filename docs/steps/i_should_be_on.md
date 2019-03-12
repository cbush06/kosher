---
layout: default
title: "I should be on the \"__\" the page"
parent: Step Definitions
---

# I should be on the "\_\_" page

Verifies the browser is on the specified page.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should be on the "([^"]*)" page$
```

## Parameters

| Position | Description | Value Type | Restrictions                     |
| :------: | ----------- | ---------- | -------------------------------- |
|    1     | page        | name       | Must be defined in `pages.json`. |

## Examples

```gherkin
Given I am on the "home" page
When I click the "Contact" link
Then I should be on the "contactUs" page
```
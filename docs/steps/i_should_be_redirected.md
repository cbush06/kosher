---
layout: default
title: "I should be redirected to the \"__\" page"
parent: Step Definitions
---

# I should be redirected to the "\_\_" page

Verifies the browser is on the specified page.
{: .fs-6 .fw-300 }

## Pattern

```
^I should be redirected to the "([^"]*)" page$
```

## Parameters

| Position | Description | Value Type | Restrictions                     |
| :------: | ----------- | ---------- | -------------------------------- |
|    1     | page        | name       | Must be defined in `pages.json`. |

## Examples

```gherkin
Given I am on the "home" page
When I click the "Contact" link
Then I should be redirected to the "contactUs" page
```
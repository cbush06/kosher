---
layout: default
title: "I am on the \"__\" page"
parent: Step Definitions
---

# I am on the "\_\_" page

Navigates the browser to the specified page.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )am on the "([^"]*)" page$
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
---
layout: default
title: "I go to the \"__\" page"
parent: Step Definitions
---

# I go to the "\_\_" page

Navigates the browser to the specified page.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )go to the "([^"]*)" page$
```

## Parameters

| Position | Description | Value Type | Restrictions                     |
| :------: | ----------- | ---------- | -------------------------------- |
|    1     | page        | name       | Must be defined in `pages.json`. |

## Examples

```gherkin
Given I go to the "home" page
When I click the "Contact" link
Then I should be on the "contactUs" page
```
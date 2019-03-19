---
layout: default
title: "I should not see a link that points to \"__\""
parent: Step Definitions
---

# I should not see a link that points to "\_\_"

Verifies that the page does not contain a link to the specified page. **Note** that kosher will attempt to find a link with just the specified parameter and with the environment URL combined with the parameter.
{: .fs-6 .fw-300 }

## Pattern

```
^I should not see a link that points to "([^"]*)"$	
```

## Parameters

| Position | Description | Value Type | Restrictions                                                                                                                               |
| :------: | ----------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
|    1     | text        | string     | Resource URL that no link on the page should point to. This needn't include the environment URL specified in the `environments.json` file. |

## Examples

```gherkin
Given I am on the "home" page
Then I should not see a link that points to "/admin"
```
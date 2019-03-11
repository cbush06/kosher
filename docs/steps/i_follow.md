---
layout: default
title: "I follow \"__\""
parent: Step Definitions
---

# I follow "\_\_"

Clicks the first hyperlink matched by the parameter.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )go to the "([^"]*)" page$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                       |
| :------: | ----------- | ---------------------------- | ---------------------------------------------------------------------------------- |
|    1     | element     | field id/name/label/selector | Must refer to a [hyperlink]({{site.baseurl}}/field_types.html#hyperlink-elements). |

## Examples

```gherkin
Given I go to the "home" page
When I follow "Contact Us"
Then I should be on the "contactUs" page
```
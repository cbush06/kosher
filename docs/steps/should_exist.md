---
layout: default
title: "the \"__\" element should exist"
parent: Step Definitions
---

# the "\_\_" element should exist

Verifies that the page contains an element matched by the selector, ID, name, or label.
{: .fs-6 .fw-300 }

## Pattern

```
^(?|the )"([^"]*)" (?|element )should exist$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | field       | field id/name/label/selector |              |

## Examples

```gherkin
Given I am on the "register" page
Then the "first_name" element should exist
```
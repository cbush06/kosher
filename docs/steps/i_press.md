---
layout: default
title: I press "__"
parent: Step Definitions
---

# I press "\_\_"

Presses/clicks first element matched by parameter 1.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )(?:press|click) "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | element     | field id/name/label/selector |              |

## Examples

```gherkin
Given I click "Contact Form"
And I press "Submit"
```
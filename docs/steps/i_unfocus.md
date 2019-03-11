---
layout: default
title: I check "__"
parent: Step Definitions
---

# I unfocus "\_\_"

Unfocuses/blurs an the specified element. First, kosher tries to build a selector for the element using its `id` or `name` attribute and uses that selector to send a JavaScript `.blur()` command to that element. If no `id` or `name` attribute is found, kosher falls back to sending the `TAB` key to the element.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )(?:unfocus|blur) "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions |
| :------: | ----------- | ---------------------------- | ------------ |
|    1     | element     | field id/name/label/selector |              |

## Examples

```gherkin
When I unfocus "First Name"
And blur "Last Name"
```
---
layout: default
title: I check "__"
parent: Step Definitions
---

# I check "__"

Set checkboxes to the <em>checked</em> state.
{: .fs-6 .fw-300 }

## Pattern

```golang
^(?:|I )check "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                     |
| :------: | ----------- | ---------------------------- | -------------------------------- |
|    1     | checkbox    | field id/name/label/selector | Must refer to a checkbox element |

## Examples

```gherkin
When I check "Accept terms"
```
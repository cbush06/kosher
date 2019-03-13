---
layout: default
title: "I send key \"__\" to \"__\""
parent: Step Definitions
---

# I send key "\_\_" to "\_\_"

Simulate a user typing a character or special key into/onto an element of the screen.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )send key "([^"]*)" to "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                 |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | key         | string                       | Must be either a single character or a special key code (see [key codes](#key-codes) below). |
|    2     | field       | field id/name/label/selector |                                                                                              |

## Examples

```gherkin
Given I am on the "input-form" page
When I send key "${ENTER}" to "comment"
And I key "Hello, World" in the "comment" field
And I send key "${BACKSPACE}" to "comment"
And I send key "${ENTER}" to "Send"
```

## Key Codes

This step will accept any of the following key codes:

| Key Code     | Meaning   | Unicode Value |
| ------------ | --------- | ------------- |
| ${BACKSPACE} | Backspace | \uE003        |
| ${ENTER}     | Enter     | \uE007        |
| ${ESCAPE}    | Escape    | \uE00C        |
| ${SPACE}     | Spacebar  | \uE00D        |
| ${DELETE}    | Delete    | \uE017        |
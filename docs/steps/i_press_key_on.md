---
layout: default
title: "I press the \"__\" on \"__\""
parent: Step Definitions
---

# I press the "\_\_" key on "\_\_"

Simulate a user typing a character or special key into/onto an element of the screen.
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )press (?:the) "([^"]*)" key (on|in) "([^"]*)"$
```

## Parameters

| Position | Description | Value Type                   | Restrictions                                                                                 |
| :------: | ----------- | ---------------------------- | -------------------------------------------------------------------------------------------- |
|    1     | key         | string                       | Must be either a single character or a special key code (see [key codes](#key-codes) below). |
|    2     | field       | field id/name/label/selector |                                                                                              |

## Examples

```gherkin
Given I am on the "input-form" page
When I press the "${ENTER}" key on "comment"
And I key "Hello, World" in the "comment" field
And I press "${BACKSPACE}" key in "comment"
And I press the "${ENTER}" key on "Send"
```

## Key Codes

This step will accept any of the following key codes. They are case-insensitive (i.e. you may use either UPPERCASE or lowercase).

| Key Code     | Meaning   | Unicode Value |
| ------------ | --------- | ------------- |
| ${BACKSPACE} | Backspace | \uE003        |
| ${ENTER}     | Enter     | \uE007        |
| ${ESCAPE}    | Escape    | \uE00C        |
| ${SPACE}     | Spacebar  | \uE00D        |
| ${DELETE}    | Delete    | \uE017        |
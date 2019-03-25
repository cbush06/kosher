---
layout: default
title: Macros 
nav_order: 44
---

# Variables
{: .no_toc }

Variables are placeholders that kosher replaces at runtime with their corresponding values.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

## What are variables?

Variables are placeholders that start with `${` followed by the variable name and ended with `}`. kosher will replace the variable with its corresponding value at runtime. Variables are used for values that cannot be typed out (e.g. special key codes) or values that are subject to change (e.g. file paths).

## Where do variables go?

Variables can be used in most value parameters of kosher steps.

## Example of Using Variables
The example below will replace `${RESOURCESDIR}` with the file path to the kosher project's `resources/` directory.

```gherkin
Given I am on the "file-upload" page
And I fill "file-upload" with "${RESOURCESDIR}/Kosher.png"
```

## Available Variables

| Variable Name | Example Usage                                              | Description                                                                          |
| ------------- | ---------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| BACKSPACE     | And I press "${BACKSPACE}" key in "comment"                | Sends a **backspace** key press to an element.                                       |
| ENTER         | And I press the "${ENTER}" key on "Send"                   | Sends a **enter** key press to an element.                                           |
| ESCAPE        | And I press the "${ESCAPE}" key on "Send"                  | Sends an **escape** key press to an element.                                         |
| SPACE         | And I press the "${SPACE}" key on "Send"                   | Sends a **space** key press to an element.                                           |
| DELETE        | And I press the "${DELETE}" key on "Send"                  | Sends a **delete** key press to an element.                                          |
| RESOURCESDIR  | And I fill "file-upload" with "${RESOURCESDIR}/Kosher.png" | Replaced with the absolute file path to the kosher project's `resources/` directory. |
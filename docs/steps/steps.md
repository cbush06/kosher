---
layout: default
title: Step Definitions
description: kosher Step Definitions.
has_children: true
has_toc: false
nav_order: 4
---

# Step Definitions

kosher provides an extensive list of step definitions for controlling the browser, interacting with page elements, and validating the application's behavior.
{: .fs-6 .fw-300 }

## List of Step Definitions

| Category  | Example                                                       | Pattern                                          |
| --------- | ------------------------------------------------------------- | ------------------------------------------------ |
| debugging | [And I wait 10 seconds](i_wait_seconds.html)                  | ^I wait (\d+) second(?:\|s)$                     |
| debugging | [And I take a screenshot](i_take_a_screenshot.html)           | ^(?:\|I ) take a screenshot$                     |
| form      | [When I fill in the following:](i_fill_in_the_following.html) | ^(?:\|I )fill in the following:$                 |
| form      | [When I fill in "Name" with "my name"](i_fill_in_with.html)   | ^(?:\|I )fill (?:\|in )"([^"]*)" with "([^"]*)"$ |
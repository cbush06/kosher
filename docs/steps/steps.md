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

| Category  | Example                                                                                                  | Pattern                                                             |
| --------- | -------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| debugging | [And I take a screenshot](i_take_a_screenshot.html)                                                      | ^(?:\|I ) take a screenshot$                                        |
| debugging | [And I wait 10 seconds](i_wait_seconds.html)                                                             | ^I wait (\d+) second(?:\|s)$                                        |
| form      | [When I fill in the following:](i_fill_in_the_following.html)                                            | ^(?:\|I )fill in the following:$                                    |
| form      | [When I fill in "Name" with "my name"](i_fill_in_with.html)                                              | ^(?:\|I )fill (?:\|in )"([^"]*)" with "([^"]*)"$                    |
| form      | [When I key in "22031" in the "Location" field](i_key_in.html)                                           | ^(?:\|I )key (?:\|in )"([^"]*)" in (?:\|the )"([^"]*)"(?:\| field)$ |
| form      | [When I select "Male" from "Sex"](i_select_from.html)                                                    | ^(?:\|I )select "([^"]*)" from "([^"]*)"$                           |
| form      | [And I check "Accept user agrement"](i_check.html)                                                       | ^(?:\|I )check "([^"]*)"$                                           |
| form      | [When I uncheck "Send me letters"](i_uncheck.html)                                                       | ^(?:\|I )uncheck "([^"]*)"$                                         |
| form      | [When I select the following values from "Filters": Accepts Gherkin Table](i_select_following_from.html) | ^(?:\|* )select (?:\|the )following values from "([^"]*)":$         |
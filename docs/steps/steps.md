---
layout: default
title: Step Definitions
description: kosher Step Definitions.
has_children: true
has_toc: false
nav_order: 5
---

# Step Definitions

kosher provides an extensive list of step definitions for controlling the browser, interacting with page elements, and validating the application's behavior.
{: .fs-6 .fw-300 }

## List of Step Definitions

| Category          | Example                                                                                                      | Pattern                                                                                      |
| ----------------- | ------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------- |
| debugging         | [And I take a screenshot](i_take_a_screenshot.html)                                                          | ^(?:\|I ) take a screenshot$                                                                 |
| debugging         | [And I wait 10 seconds](i_wait_seconds.html)                                                                 | ^I wait (\d+) second(?:\|s)$                                                                 |
| form              | [When I fill in the following:](i_fill_in_the_following.html)                                                | ^(?:\|I )fill in the following:$                                                             |
| form              | [When I fill in "Name" with "my name"](i_fill_in_with.html)                                                  | ^(?:\|I )fill (?:\|in )"([^"]*)" with "([^"]*)"$                                             |
| form              | [When I key in "22031" in the "Location" field](i_key_in.html)                                               | ^(?:\|I )key (?:\|in )"([^"]*)" in (?:\|the )"([^"]*)"(?:\| field)$                          |
| form              | [When I select "Male" from "Sex"](i_select_from.html)                                                        | ^(?:\|I )select "([^"]*)" from "([^"]*)"$                                                    |
| form              | [And I check "Accept user agrement"](i_check.html)                                                           | ^(?:\|I )check "([^"]*)"$                                                                    |
| form              | [When I uncheck "Send me letters"](i_uncheck.html)                                                           | ^(?:\|I )uncheck "([^"]*)"$                                                                  |
| form              | [When I select the following values from "Filters": Accepts Gherkin Table](i_select_following_from.html)     | ^(?:\|* )select (?:\|the )following values from "([^"]*)":$                                  |
| form              | [When I unselect the following values from "Filters": Accepts Gherkin Table](i_unselect_following_from.html) | ^(?:\|I )unselect (?:\|the )following values from "([^"]*)":$                                |
| form              | [When I choose the "radio1"](i_choose.html)                                                                  | ^(?:\|I )choose (?:\|the )"([^"]*)" radio$                                                   |
| form              | [When I press "Submit"](i_press.html)                                                                        | ^(?:\|I )(?:press\|click) "([^"]*)"$                                                         |
| form              | [When I press the 4th instance of "nextButton"](i_press_instance.html)                                       | ^(?:\|I )(?:press\|click) the (first\|last\|[0-9]+(?:th\|st\|rd\|nd)) instance of "([^"]*)"$ |
| form              | [When I press the "Submit" button](i_press_button.html)                                                      | ^(?:\|I )(?:press\|click) the "([^"]*)" (?:button\|link)$                                    |
| form              | [When I unfocus "first_name"](i_unfocus.html)                                                                | ^(?:\|I )(?:unfocus\|blur) "([^"]*)"$                                                        |
| form              | [And I hover over "BidPlansMenu"](i_hover.html)                                                              | ^(?:\|I )hover over "([^"]*)"$                                                               |
| form              | [Then I enter today's date in "Date"](i_enter_todays_date.html)                                              | ^(?:\|I )enter today's date in "([^"]*)"$                                                    |
| form verification | [Then I verify "date" has today's date](i_verify_todays_date.html)                                           | ^(?:\|I )verify "([^"]*)" has today's date$                                                  |
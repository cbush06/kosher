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

| Category          | Example                                                                                                      | Element Type                        |
| ----------------- | ------------------------------------------------------------------------------------------------------------ | ----------------------------------- |
| debugging         | [And I take a screenshot](i_take_a_screenshot.html)                                                          | N/A                                 |
| debugging         | [And I wait 10 seconds](i_wait_seconds.html)                                                                 | N/A                                 |
| form              | [When I fill in the following:](i_fill_in_the_following.html)                                                | text-based, radio, checkbox, select |
| form              | [When I fill in "Name" with "my name"](i_fill_in_with.html)                                                  | text-based, radio, checkbox, select |
| form              | [When I key in "22031" in the "Location" field](i_key_in.html)                                               | text-based                          |
| form              | [When I select "Male" from "Sex"](i_select_from.html)                                                        | select                              |
| form              | [And I check "Accept user agrement"](i_check.html)                                                           | checkbox                            |
| form              | [When I uncheck "Send me letters"](i_uncheck.html)                                                           | checkbox                            |
| form              | [When I select the following values from "Filters": Accepts Gherkin Table](i_select_following_from.html)     | select                              |
| form              | [When I unselect the following values from "Filters": Accepts Gherkin Table](i_unselect_following_from.html) | select                              |
| form              | [When I choose the "radio1"](i_choose.html)                                                                  | radio                               |
| form              | [When I press "Submit"](i_press.html)                                                                        | any                                 |
| form              | [When I press the 4th instance of "nextButton"](i_press_instance.html)                                       | any                                 |
| form              | [When I press the "Submit" button](i_press_button.html)                                                      | button, hyperlink                   |
| form              | [When I unfocus "first_name"](i_unfocus.html)                                                                | any                                 |
| form              | [And I hover over "BidPlansMenu"](i_hover.html)                                                              | any                                 |
| form              | [Then I enter today's date in "Date"](i_enter_todays_date.html)                                              | text-based                          |
| form verification | [Then I verify "date" has today's date](i_verify_todays_date.html)                                           | text-based                          |
| form verification | [Then "date" should contain today's date](should_contain_todays_date.md)                                     | text-based                          |
| form verification | [Then "Name" should contain "my name"](should_contain.html)                                                  | text-based                          |
| form verification | [Then "Name" should not contain "my name"](should_not_contain.html)                                          | text-based                          |
| form verification | [Then "states" should have the following options:](should_have_options.html)                                 | select                              |
| form verification | [Then "states" should have the following options selected:](should_have_options_selected.html)               | select                              |
| form verification | [And "states" should not have the following options selected:](should_not_have_options_selected.html)        | select                              |
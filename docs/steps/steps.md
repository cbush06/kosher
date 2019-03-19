---
layout: default
title: Step Definitions
description: kosher Step Definitions.
has_children: true
has_toc: false
nav_order: 40
---

# Step Definitions

kosher provides an extensive list of step definitions for controlling the browser, interacting with page elements, and validating the application's behavior.
{: .fs-6 .fw-300 }

## List of Step Definitions

| Category                | Example                                                                                                      | Element Type                                |
| ----------------------- | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------- |
| debugging               | [When I take a screenshot](i_take_a_screenshot.html)                                                         | N/A                                         |
| debugging               | [When I wait 10 seconds](i_wait_seconds.html)                                                                | N/A                                         |
| form                    | [When I fill in the following:](i_fill_in_the_following.html)                                                | text-based, radio, checkbox, select         |
| form                    | [When I fill in "Name" with "my name"](i_fill_in_with.html)                                                  | text-based, radio, checkbox, select         |
| form                    | [When I fill in "Name" with:](i_fill_in_with_multiline.html)                                                 | text-based                                  |
| form                    | [When I key in "22031" in the "Location" field](i_key_in.html)                                               | text-based                                  |
| form                    | [When I select "Male" from "Sex"](i_select_from.html)                                                        | select                                      |
| form                    | [When I check "Accept user agrement"](i_check.html)                                                          | checkbox                                    |
| form                    | [When I uncheck "Send me letters"](i_uncheck.html)                                                           | checkbox                                    |
| form                    | [When I select the following values from "Filters": Accepts Gherkin Table](i_select_following_from.html)     | select                                      |
| form                    | [When I unselect the following values from "Filters": Accepts Gherkin Table](i_unselect_following_from.html) | select                                      |
| form                    | [When I choose the "radio1" radio](i_choose.html)                                                            | radio                                       |
| form                    | [When I press/click "Submit"](i_press.html)                                                                  | any                                         |
| form                    | [When I press/click the 4th instance of "nextButton"](i_press_instance.html)                                 | any                                         |
| form                    | [When I press/click the "Submit" button/link](i_press_button.html)                                           | button, hyperlink                           |
| form                    | [When I unfocus/blur "first_name"](i_unfocus.html)                                                           | any                                         |
| form                    | [When I hover over "BidPlansMenu"](i_hover.html)                                                             | any                                         |
| form                    | [Then I enter today's date in "Date"](i_enter_todays_date.html)                                              | text-based                                  |
| form verification       | [Then I verify "date" has today's date](i_verify_todays_date.html)                                           | text-based                                  |
| form verification       | [Then "date" should contain today's date](should_contain_todays_date.md)                                     | text-based                                  |
| form verification       | [Then "Name" should contain "my name"](should_contain.html)                                                  | text-based                                  |
| form verification       | [Then "Name" should contain:](should_contain_multiline.html)                                                 | text-based                                  |
| form verification       | [Then "Name" should not contain "my name"](should_not_contain.html)                                          | text-based                                  |
| form verification       | [Then "Name" should not contain:](should_not_contain_multiline.html)                                         | text-based                                  |
| form verification       | [Then "states" should have the following options:](should_have_options.html)                                 | select                                      |
| form verification       | [Then "states" should have the following options selected:](should_have_options_selected.html)               | select                                      |
| form verification       | [Then "states" should not have the following options selected:](should_not_have_options_selected.html)       | select                                      |
| form verification       | [Then the "Accept user agrement" checkbox should be checked](checkbox_should_be_checked.html)                | checkbox                                    |
| form verification       | [Then the "Do not touch me" checkbox should not be checked](checkbox_should_not_be_checked.html)             | checkbox                                    |
| form verification       | [Then the "radio 1" radio should be selected](radio_should_be_selected.html)                                 | radio                                       |
| form verification       | [Then the "radio 2" radio should not be selected](radio_should_not_be_selected.html)                         | radio                                       |
| navigation              | [Given I am on the "home" page](i_am_on_page.html)                                                           | N/A                                         |
| navigation              | [When I go to the "other" page](i_go_to_page.html)                                                           | N/A                                         |
| navigation              | [When I follow "Privacy Policy"](i_follow.html)                                                              | N/A                                         |
| navigation verification | [Then I should be redirected to the "congratulations" page](i_should_be_redirected.html)                     | N/A                                         |
| navigation verification | [Then I should be on the "congratulations" page](i_should_be_on.html)                                        | N/A                                         |
| browser                 | [When I maximize the window](i_maximize.html)                                                                | N/A                                         |
| browser                 | [When I switch to the first/last window](i_switch_to_window.html)                                            | N/A                                         |
| browser                 | [When I switch to frame 2](i_switch_frame_num.html)                                                          | frame                                       |
| browser                 | [When I accept the popup](i_accept_popup.html)                                                               | N/A                                         |
| browser                 | [When I decline the popup](i_decline_popup.html)                                                             | N/A                                         |
| browser                 | [When I switch to iframe 2](i_switch_iframe_num.html)                                                        | iframe                                      |
| browser                 | [When I switch to the root frame](i_switch_root_frame.html)                                                  | N/A                                         |
| browser verification    | [Then I should see the popup text "Hello, World!"](i_should_see_popup_text.html)                             | N/A                                         |
| browser verification    | [And I should not see the popup text "Hello, World!"](i_should_not_see_popup_text.html)                      | N/A                                         |
| verification            | [And I should see "Great, you can click links!"](i_should_see.html)                                          | N/A                                         |
| verification            | [And I should not see "some bla-bla"](i_should_not_see.html)                                                 | N/A                                         |
| verification            | [Then I should see all of the texts:](i_should_see_all_of.html)                                              | N/A                                         |
| verification            | [Then I should see the following list:](i_should_see_the_following.html)                                     | N/A                                         |
| verification            | [Then I should see the "Submit" button/link](i_should_see_button_link.html)                                  | N/A                                         |
| verification            | [And I should not see the "Click Me" button/link](i_should_not_see_button_link.html)                         | N/A                                         |
| verification            | [Then I should see a link that points to "/about-us"](i_should_see_url_link.html)                            | N/A                                         |
| verification            | [And I should not see a link that points to "/contact-us"](i_should_not_see_url_link.html)                   | N/A                                         |
| verification            | [Then "Submit" should be disabled](should_be_disabled.html)                                                  | text-based, radio, checkbox, select, button |
| verification            | [Then "Submit" should be enabled](should_be_enabled.html)                                                    | text-based, radio, checkbox, select, button |
| verification            | [Then the 5th instance of "Page1PaginationLinks" should be disabled](nth_instance_disabled.html)             | text-based, radio, checkbox, select, button |
| verification            | [Then the 5th instance of "Page1PaginationLinks" should be enabled](nth_instance_enabled.html)               | text-based, radio, checkbox, select, button |
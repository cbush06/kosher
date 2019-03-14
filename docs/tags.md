---
layout: default
title: Tags
nav_order: 9
---

# Tags
{: .no_toc }

Gherkin tags are an easy method for running subsets of your overall test suite.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

## What are tags?

Tags are a simple way of organizing/categorizing features, scenarios, scenario outlines, and examples. Tags are any space-delimited 
string in a feature file that begins with the 'at' (@) symbol. After you've added tags to your `*.feature` files, you can specify
which tags to execute when you run the tests.

## Where can tags go?

Kosher (via [GoDog](https://github.com/DATA-DOG/godog)) supports tagging features, scenarios, scenario outlines, and examples. 

## Example `*.feature` File
```gherkin
@Navigation
Feature: Verify Results of Navigation Actions

    After clicking a link, clicking a button, or submitting
    a form, verfiy that the correct navigations take place.

    Background: Start on the Table Filter page
        Given I go to the "table-search" page
        And I maximize the window

    @Smoke
    Scenario: Verify Page
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be on the "bootstrap-date-picker" page

    @Regression
    Scenario: Verify Page with Trailing Slash
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Demo Home" link
        Then I should be on the "home" page

    @Smoke @Regression
    Scenario Outline: Verify Redirect
        After clicking on a link, verify redirection to the appropriate page.

        Given I click the "<link1_selector>" link
        And I follow the "<link2_selector>" link
        Then I should be redirected to the "<page_name>" page

        @FormsNavigation
        Examples:
        | link1_selector | link2_selector     | page_name      |
        | Input Forms    | Simple Form Demo   | basicform      |
        | Input Forms    | Checkbox Demo      | basic-checkbox |
        | Input Forms    | Radio Buttons Demo | radio-button   |

        @TablesNavigation
        Examples:
        | link1_selector | link2_selector      | page_name    |
        | Table          | Table Data Search   | table-search |
        | Table          | Table Sort & Search | table-sort   |
```

## How to Run Specific Tags

In the example above, executing `kosher run -t @Navigation` would run everything in the _Verify Results of Navigation Actions_ feature. 

If, instead, you executed `kosher run -t @Smoke`, only the `Background` section, _Verify Page_ scenario, and _Verify Redirect_ scenario outline
would be ran. Likewise, executing `kosher run -t @Regression` would cause the `Background` section, _Verify Page with Trailing Slash_ scenario,
and _Verify Redirect_ scenario outline to be ran.

To run multiple scenarios from the feature file, you could add the `@Regression` tag like this: `kosher run -t @Smoke,@Regression`. This command

Finally, you can also run specific example lists for scenario outlines by specifying tags: `kosher run -t @FormsNavigation`
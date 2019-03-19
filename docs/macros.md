---
layout: default
title: Tags
nav_order: 43
---

# Tags
{: .no_toc }

Macros are reusable collections of steps and help solve the problem of duplication in kosher tests.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

## What are macros?

Macros are named collections of steps (henceforth, called *substeps*) that become steps you can use in your test scripts. Executing a macro step executes all its *substeps*.

## Where do macros go?

Macros are stored in the `macros/` directory of a kosher project.

## How do I create a macro?

1. Create a new `*.feature` file in the `macros/` directory.
2. You may enter whatever title and description you want in the `Feature:` block (this block does not affect your macros).
    ```gherkin
    Feature: Form Filling Macros
        Macros related to filling an input form.
    ```
3. Create scenarios with titles that will become your macro step definitions. You can create many scenarios in a single macro `*.feature` file. All of the scenarios will become available as macro steps in your test scripts.

    For the example below, you could use the macro below in a test script like this `When I fill the input form`. Doing so would execute all the steps in the macro's scenario.
    ```gherkin
    Scenario: I fill the input form
        Fill the Selenium Easy input form demo.

        Then I fill in "first_name" with "Gherkin"
        And I fill in "last_name" with "Warrior"
    ```
4. Use the macros' *scenario titles* as steps in your test scripts. The example below provides a full demonstration.

## Example Macro `*.feature` File in `macros/` Directory
```gherkin
Feature: Form Filling Macros
    Macros related to filling forms.

    Scenario: I fill the input form
        Fill the Selenium Easy input form demo.

        Then I fill in "first_name" with "John"
        And I fill in "last_name" with "Doe"

    Scenario: I fill the rest of the input form
        Fill the other fields in the Selenium Easy input form demo.

        Then I fill in "email" with "jdoe@gmail.com"
        And I fill in "phone" with "(444)333-2222"
        And I fill in "address" with "222 Ausley Road"
        And I fill in "city" with "Albany"
        And I select "New York" from "state"
        And I fill in "zip" with "333333"
        And I choose the "Yes" radio
```

## Example Test Script `*.feature` File Using the Macros Above
```gherkin
Feature: Test Macros
    Verify that Kosher's macros feature
    works correctly.

    Background:
        Given I maximize the window

    Scenario: Fill a Textbox
        Verify steps for fill a textbox.

        Given I go to the "input-form" page
        When I fill the input form
        Then "first_name" should contain "John"
        And "last_name" should contain "Doe"
        And I fill the rest of the input form
        Then the "Yes" radio should be selected
```

## Limitations of Macros

The following do not work in macros:

* Tags
* DocStrings
* DataTables
* Scenario Outlines
* Examples
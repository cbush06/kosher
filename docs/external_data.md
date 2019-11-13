---
layout: default
title: External Data 
nav_order: 44
---

# External Data
{: .no_toc }

Kosher can use CSV data files to supply the `Examples` data to a `Scenario Outline` step.
{: .fs-6 .fw-300 }

Often times, testers will want to use external sources to supply their test data. Kosher supports this by using CSV files to populate the `Examples` table
of a `Scenario Outline` step.

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

## Where does the CSV go?
Kosher will look for CSV files in the project's `resources/` directory.

## How do I reference the CSV file for a `Scenario Outline`?
To populate a `Scenario Outline` step with data from CSV file, you add a special `@loadcsv(filename.csv)` tag above the `Scenario Outline`. The text between
the parentheses of this tag must exactly match the filename of a CSV file in the project's `resources/` directory.

## How should the `Examples` table of the `Scenario Outline` be formatted?
You must provide column names and 1 blank row. This row will be removed by kosher, but is required by kosher's Gherkin parser.

> **Note:** You do not need to list every column name that exists in the CSV file in your empty `Examples` table, but the column names that you do list **must**
match their corresponding column name (case sensitive) in the CSV file.

## How should the CSV file be formatted?
Kosher interprets the first row of the CSV file as column names. The column names you list in the empty `Examples` table of a `Scenario Outline` must match their
corresponding column names in the CSV file; however, you do not have to list all of the column names that exist in the CSV file in the `Scenario Outline`.

## @loadcsv(...) Example
The example that follows shows how you might use a CSV file to populate the `Examples` table of a `Scenario Outline`.

### Sample CSV Data File
```csv
FirstName,LastName,Email,City,State,Phone
John,Doe,jdoe@gmail.com,New York,NY,3934445555
Jane,Doe,jdoe2@gmail.com,Boston,MA,2223334444
Babe,Ruth,bruth@gmail.com,Miami,FL,5556667777
```

### Sample `Scenario Outline` Populated by CSV Data File
```gherkin
@loadcsv(test.csv)
    Scenario Outline: Verify Scenario Outline
        Given I fill in "<Email>" 
        And I click the "OK" button
        And I fill in "<Phone>"
        And I click the "OK" button

        Examples:
            | Email | Phone |
            |       |       |

@loadcsv(LogwayApprovalsTest.csv)
    Scenario Outline: Verify requested user has completed Self-Registration
        Given I fill in "frmManageUsers:tblUsers:txtFilterLastName" with "<Last Name>"
        And I fill in "frmManageUsers:tblUsers:txtFilterFirstName" with "<First Name>"
        And I wait 2 seconds
        And I should see "<UserEmailAddress>"

        Examples:
            | Last Name | First Name | UserEmailAddress |
            |           |            |                  |
```
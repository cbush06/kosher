---
layout: default
title: "I should the following list/errors:"
parent: Step Definitions
---

# I should see the following list/errors:

Verifies that the page contains all of the specified texts. This step is simply an alias of `I should see all of the texts:` that accepts multiple columns of texts. **Note** that this step may sometimes fail if any of the texts are split across multiple elements (e.g. parts of a text are styled individually using multiple `<span></span>` elements).
{: .fs-6 .fw-300 }

## Pattern

```
^(?:|I )should see the following(?:| errors| list):$
```

## Parameters

None.

## Data Table Columns

This step's data table may have multiple columns. Ultimately, the data table will be processed
one row at the time, one cell at the time. The page will be searched for the contents of each cell
and this step will fail on the first cell contents that could not be found.

Each cell may include [variables](../variables.html).

## Examples

```gherkin
Given I am on the "home" page
Then I should see the following list:
    | # | Task         | Assignee   | Status      |
    | 1 | Wireframes   | John Smith | in progress |
    | 2 | Landing Page | Mike Trout | completed   |
    | 3 | SEO tags     | Loblab Dan | failed qa   |
```

**Note** that the first row above _IS NOT_ interpreted as a table header. It is simply searched for in the page like all the other rows.

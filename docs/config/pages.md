---
layout: default
title: pages.json
description: Format of pages.json
parent: Configuration
nav_order: 2
---

# pages.json

The `pages.json` file allows the user to define page names to be used in Gherkin steps along with the pages' corresponding URLs. These URLs should only include the portion of the web address that comes after the URLs defined in `environments.json`.
{: .fs-6 .fw-300 }

## Format

```json
{
    "home": "/",
    "basicform": "basic-first-form-demo.html",
    "table-sort": "table-sort-search-demo.html"
}
```

## Relation to environments.json

As mentiond before, `environments.json` contains a base URL of application being tested for each of its environments.

The `pages.json` file contains the portion of each page's URL that follows the environment URLs. The names assigned to each page URL are used in the Gherkin steps for navigating to different pages and verifying that redirections and navigations were performed correctly.

For example, given an `environments.json` file with the following content:

```json
{
    "test": "https://about.google"
}
```

and a `pages.json` file with this content:

```json
{
    "prods": "/products"
}
```

The Gherkin step `When I go to the "prods" page` would cause the browser to navigate to `https://about.google/products`.
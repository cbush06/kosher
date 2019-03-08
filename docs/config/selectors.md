---
layout: default
title: selectors.json
description: Format of selectors.json
parent: Configuration
nav_order: 3
---

# selectors.json

The `selectors.json` file allows the user to define CSS or XPath selectors to use to specify which elements of pages/screens to interact with or validate.
{: .fs-6 .fw-300 }

## Format

```json
{
	"header1Css": "css:html > body > table > tbody > tr > th:first-child",
	"header2Xpath": "xpath:/html/body/table/tbody/tr[1]/th[2]"
}
```

Each selector must have a unique name. Selector values must begin with either `css:` or `xpath:` to identify which form of selector is being defined.
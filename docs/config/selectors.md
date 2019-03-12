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
	"header1": "css: html > body > table > thead > tr > th:first-child",
	"header2": "xpath: /html/body/table/tbody/tr[1]/th[2]",
    "row1Cell3": "css: html > body > table > tbody > tr > td:nth-child(3)"
}
```

Each selector must have a unique name. Selector values must begin with either `css:` or `xpath:` to identify which form of selector is being defined.

## Using Selectors in Step Defitions

Given the example above and this web page you might be testing:

```html
<html>
    <head>
        <title>Some Web Page</title>
    </head>
    <body>
        <table>
            <thead>
                <tr>
                    <th>Product ID</th>
                    <th>Product Name</th>
                    <th>Product Price</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>1399648</td>
                    <td>Golang for Dummies</td>
                    <td>$39.98</td>
                </tr>
            </tbody>
        </table>
    </body>
</html>
```

The first and second steps below will pass and the third will fail:

```gherkin
Then "header1Css" should contain "Product ID"
And "header2Xpath" should contain "Product Name"
And "row1Cell3" should contain "$40.00"
```
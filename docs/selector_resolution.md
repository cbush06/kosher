---
layout: default
title: Selector Resolution
nav_order: 8
---

# Selector Resolution
{: .no_toc }

kosher follows a very specific algorithm when attempting to identify an element using either a literal selector or one defined in the `selectors.json` file.
{: .fs-6 .fw-300 }

1. TOC
{:toc}

## Types of Selectors

### Literal Selectors

A **literal selector** is one that is not defined in the `selectors.json` file.

For example, `First Name` in the script below is a literal selector because it is not defind in the `selecors.json` file and, presumably, the intention is to match a textbox with an associated label of `First Name`.

```gherkin
When I fill in "First Name" with "John"
```

### Defined Selectors

A **defined selector** is one that is defined in the `selectors.json` file and may be either a CSS or an XPath selector. For more information on CSS and XPath selectors, see [selectors.json]({{site.baseurl}}/config/selectors.html).

For example, assuming there is an HTML element `<input type="text" id="firstName" />` and the `selectors.json` file is defined as:

```json
{
    "firstNameTextbox": "css: #firstName"
}
```

Then, the following would be an example of using a defined selector:

```gherkin
When I fill in "firstNameTextbox" with "John"
```

## Selector Algorithms

### Literal Selectors

Literal selectors follow a specific algorithm built into kosher. This algorithm builds a **list of matches** by following these steps:
1. Find all form fields with associated **labels** that match the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list
2. Find all **buttons** with texts that match the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list
3. Find all **hyperlinks** with texts that match the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list
4. Find all form elements with **`name` attributes** that match the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list
5. Find all form elements with **`id` attributes** that match the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list
6. Find any element on the page that has **text** content that matchs the literal selector
   1. If the `ignoreInvisible` setting is true, add only visible matches to the list;
   2. Otherwise, add all matches to the list

This algorithm is applied in all steps that do not have a narrower scope. To determine the scope of a step's selector resolution, see the **Element Type** column on the [Step Definitions]({{site.baseurl}}/steps/steps.html) page.

Unless otherwise stated in a step's documentation, the first element in the list will be acted upon by the step.

### Defined Selectors

Defined selectors are fairly straightforward; they are resolved using a browser's native CSS or XPath utilities.
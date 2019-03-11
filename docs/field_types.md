---
layout: default
title: Field Types
description: kosher Supported Field Types
nav_order: 4
---

# Supported Field Types
{: .no_toc }

kosher supports most common HTML field types.
{: .fs-6 .fw-300 }

## Field Types
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Buttons and Hyperlinks

Most steps that support clicking/pressing buttons also support hyperlinks. Additionally, unless otherwise stated, these steps support pressing/clicking any element visible on the screen.

### Button Elements

```html
<input type="button" />
<input type="reset" />
<input type="submit" />
<button></button>
```

### Hyperlink Elements

```html
<a href="#"></a>
```

## Text-based Fields

kosher supports interacting with all text-based HTML fields. Examples of those follow:

```html
<input type="date" />
<input type="datetime-local" />
<input type="email" />
<input type="month" />
<input type="number" />
<input type="password" />
<input type="search" />
<input type="tel" />
<input type="time" />
<input type="url" />
<input type="week" />
<textarea></textarea>
<input type="text" />
```

## Checkbox Fields

Some steps allow checking or un-checking checkbox fields. For the purposes of kosher, checkbox fields must actually be an HTML checkbox element:

```html
<input type="checkbox" name="favorite_colors_green" />
<input type="checkbox" name="favorite_colors_red" />
```

## Radio Button Fields

kosher allows selecting a value from a group of radio buttons linked together by a common name value. Again, with kosher, the radio button fields must actually be an HTML radio button element:

```html
<input type="radio" name="favorite_colors" value="green" />
<input type="radio" name="favorite_colors" value="red" />
```

## Select Fields

Some step definitions enable interactions with select lists. As with the above types, these must actuall be HTML select elements. Pay _careful attention_ to the documentation for these step definitions, as some restrict their applicability to select lists with the `multiple` attribute set.

```html
<select name="state">
    <option value="al">Alabama</option>
    <option value="ak">Alaska</option>
    <option value="ar">Arkansas</option>
</select>

<select name="favorite_colors" size="2" multiple>
    <option value="green">Green</option>
    <option value="red">Red</option>
</select>
```

## Browser Pop-ups

kosher also provides capabilities for manipulating browser prompts (alerts, confirmations, and prompts). These type of pop-ups are generated via JavaScript commands such as the following:

```javascript
window.alert('Hello, world!');

let doesAccept = window.confirm('Do you accept?');

let firstName = window.prompt('Enter your first name:');
```
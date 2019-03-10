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
## Selection Fields

Some step definitions enable interactions with the 3 selection field types: checkboxes, radio buttons, and selects. Examples of these follow:

```html
<input type="checkbox" name="favorite_colors_green" />
<input type="checkbox" name="favorite_colors_red" />

<input type="radio" name="favorite_colors" value="green" />
<input type="radio" name="favorite_colors" value="red" />

<select name="favorite_colors">
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
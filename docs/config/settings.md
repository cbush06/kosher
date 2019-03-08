---
layout: default
title: settings.json
description: Options for settings.json
parent: Configuration
nav_order: 4
---

# settings.json

The `settings.json` file allows the user to specify global settings for a kosher project. Some values in the file are defaults but can be overridden by flags of the `run` command (e.g. `--appVersion`, `--environment`).
{: .fs-6 .fw-300 }

## Format

```json
{
    "projectName": "kosher",
    "appVersion": "1.0.0",
    "platform": "web",
    "driver": "chrome",
    "reportFormat": "pretty",
    "dateFormat": "MM/DD/YYYY",
    "defaultEnvironment": "test",
    "screenFormat": "desktop",
    "quitOnFail": false,
    "screenFormats": {
        "desktop": {
            "width": 2000,
            "height": 980
         },
        "mobile": {
            "width": 362,
            "height": 868
        },
        "tablet": {
            "width": 814,
            "height": 868
        },
        "landscape": {
            "width": 522,
            "height": 362
        }
    }
}
```

## Properties

### projectName

Sets the name of the project recorded in results files.

| Characteristic | Description       |
| -------------- | ----------------- |
| type           | string            |
| required       | no                |
| default        | kosher tested app |

### appVersion

Sets the version of the project recorded in results files.

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | 1.0.0       |

### platform

Specifies if steps are tailored for desktop or web use. Valid options are: `web`, `desktop`

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | web         |

### driver

Specifies which webdriver is used to execute the tests. Valid options are `chrome`, `ie`, `phantomjs`, `desktop`

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | chrome      |

### reportFormat

Specifies what format is used to report test results. Valid options are: `pretty`, `html`, `bootstrap`, `simple`, `pretty`, `progress`, `junit`, `cucumber`

For descriptions of the reports, see [Reports](../reports/reports.html)

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | pretty      |

### dateFormat

Specifies the format used when validating or setting dates.

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | MM/DD/YYYY  |

Valid placeholders are:

| Placeholder | Description   | Example |
| ----------- | ------------- | ------- |
| MMMM        | Full month    | January |
| MMM         | Short month   | Jan     |
| MM          | Numeric month | 01      |
| YYYY        | Full year     | 2006    |
| YY          | Short year    | 06      |
| DDDD        | Full day      | Monday  |
| DDD         | Short day     | Mon     |
| DD          | Numeric day   | 02      |

### defaultEnvironment

Specifies the default environment to execute tests in. For more information, see [environments.json](environments.html).

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | test        |

### screenFormat

Specifies which dimensions to apply to newly opened window and when the `I maximize the window` step is executed. For more information, see [screenFormats](#screenformats) below.

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | desktop     |

### quitOnFail

Specifies if kosher continues executing remaining scenarios after a step fails.

| Characteristic | Description |
| -------------- | ----------- |
| type           | boolean     |
| required       | no          |
| default        | false       |

### screenFormats

Defines available screen formats. One of these must be selected by the [screenFormat](#screenformat) option above.

Each option in the `screenFormats` object must define a `width` and `height` property, both of type `int`.

<table>
    <thead>
        <tr>
            <th>Characteristic</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>type</td>
            <td>object</td>
        </tr>
        <tr>
            <td>required</td>
            <td>no</td>
        </tr>
        <tr>
            <td>default</td>
            <td><pre>
{
    "desktop": {
        "width": 2000,
        "height": 980
    },
    "mobile": {
        "width": 362,
        "height": 868
    },
    "tablet": {
        "width": 814,
        "height": 868
    },
    "landscape": {
        "width": 522,
        "height": 362
    }
}</pre>
            </td>
        </tr>
    </tbody>
</table>
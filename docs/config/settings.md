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
    "driver": "chrome",
    "reportFormat": "pretty",
    "dateFormat": "MM/DD/YYYY",
    "defaultEnvironment": "test",
    "screenFormat": "desktop",
    "quitOnFail": false,
    "ignoreInvisible": true,
    "waitAfterScenario": 0,
    "waitAfterStep": 0,
    "accessibility": {
		"ruleSets": [
			"wcag21aa",
			"section508"
		],
		"impactThreshold": "serious"
	},
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
    },
	"integrations": {
		"jira": {
			"host": "http://127.0.0.1",
			"defaults": {
				"projectKey": "PROJE",
				"issueType": "Bug",
				"affectsVersion": "1.0.0",
				"labels": "test,functional,kosher",
				"priority": "Normal"
			}
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

### driver

Specifies which webdriver is used to execute the tests. Valid options are `chrome`, `ie`, `phantomjs`, `desktop`

| Characteristic | Description |
| -------------- | ----------- |
| type           | string      |
| required       | no          |
| default        | chrome      |

### reportFormat

Specifies what format is used to report test results. Valid options are: `pretty`, `bootstrap`, `simple`, `progress`, `junit`, `cucumber-json`

For descriptions of the reports, see [Reports](../reports/reports.html)

| Characteristic | Description          |
| -------------- | -------------------- |
| type           | string, string array |
| required       | no                   |
| default        | pretty               |

You can use multiple formats by setting the `reportFormat` to a JSON string array: `["pretty", "bootstrap"]`

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

### ignoreInvisible

Specifies if kosher should act as if elements that are invisible to the user do not exist.

| Characteristic | Description |
| -------------- | ----------- |
| type           | boolean     |
| required       | no          |
| default        | true        |

### waitAfterScenario

Specifies a wait time to pause for following every scenario. **This is measured in milliseconds.**

| Characteristic | Description |
| -------------- | ----------- |
| type           | int         |
| required       | no          |
| default        | 0           |

### waitAfterStep

Specifies a wait time to pause for following every step. **This is measured in milliseconds.**

| Characteristic | Description |
| -------------- | ----------- |
| type           | int         |
| required       | no          |
| default        | 0           |

### accessibility

Configures how the Axe accessiblity API will scan the page for issues.

#### ruleSets

The Axe tags to apply to the page. These are effectively sets of rules. You can find a list of available rule sets at [Axe API Documentation](https://www.deque.com/axe/axe-for-web/documentation/api-documentation/#api-name-axegetrules).

| Characteristic | Description                           |
| -------------- | ------------------------------------- |
| type           | string array                          |
| required       | no                                    |
| default        | [ "wcag2a", "wcag2aa", "section508" ] |

#### impactThreshold

Specifies the impact (i.e. severity) that an accessibility finding must be to fail the [I test the page for accessibility](../steps/i_test_for_accessibility.html) step. See the [Axe API Documentation](https://www.deque.com/axe/axe-for-web/documentation/api-documentation/#results-object) for more details.

| Characteristic    | Description                        |
| ----------------- | ---------------------------------- |
| type              | string                             |
| required          | no                                 |
| default           | critical                           |
| available options | minor, moderate, serious, critical |


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
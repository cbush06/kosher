---
layout: default
title: settings.json
description: Options for settings.json
parent: Configuration
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
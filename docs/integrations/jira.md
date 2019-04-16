---
layout: default
title: Jira Integration
description: Details of kosher's Jira integration.
parent: Integrations
nav_order: 1
---

# Jira Integration
{: .no_toc }

kosher has the ability to create new Jira issues from failed test scenarios.
{: .fs-6 .fw-300 }

1. TOC
{:toc}

---

## Prerequisites

* You have a Jira account
* You have executed the kosher test(s) with a report format that produces a `results.json` file in the `resuls` directory (e.g. cucumber, simple, bootstrap, html)
* You have add Jira integration configurations to the `settings.json` file in the `config` directory

## Command-line Interface

```bash
$ kosher send jira
```

For more details, see the [CLI](../../cli.html#jira) page.

## Configuration

The Jira integration requires, at a minimum, that the host be specified. Other optional settings may, also, be specified.

### Format

```bash
{
    ...,
    "integrations": {
        "jira": {
            "host": "https://jira.myserver.com",
            "labels": "bug,kosher,failure",
            "summaryTemplate": "jira_summary.txt",
            "descriptionTemplate": "jira_description.txt"
        }
    }
}
```

### Settings

| Setting             | Description                                                                                                             |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| host                | URL of the Jira server to connect to.                                                                                   |
| labels              | Comma-delimited list of labels to add to new issues.                                                                    |
| summaryTemplate     | Name of file in the `config` directory that contains a Golang template file to be used for creating issue summaries.    |
| descriptionTemplate | Name of file in the `config` directory that contains a Golang template file to be used for creating issue descriptions. |

## Templates

kosher's Jira integration creates the summary and description using Golang templates. kosher has default templates embedded in its source, but you may create your own in files in the `config` directory and specify those files in the `settings.json` file. See [Configuration Settings](#settings) above for details on specifying template files.

For more information on Golang's templates, see [text/template](https://golang.org/pkg/text/template/).

### Default Templates

Default summary template:
```
{{.Feature.Name}}: {{.Element.Name}}
```

Default description template:
```
h2. Issue:
{{.FailedStep.Result.Error}}

h2. Feature Title:
{{.Feature.Name}}

h2. Scenario Title:
{{.Element.Name}}

h2. Scenario Description
{{.Element.Description}}

h2. Failed Step / Actual Result:
{color:red}*{{.FailedStep.GetTrimmedKeyword}}* {{.FailedStep.Name}}{color}
{color:red}{{.FailedStep.Result.Error}}

h2. Steps / Expected Result:
{{range .Element.Steps}}# *{{.GetTrimmedKeyword}}* {{.Name}}
{{end}}
```

### Template Context

kosher provides the following Golang struct as the context for Jira summary and description templates:

```golang
type IssueContext struct {
	Feature    *report.CukeFeature
	Element    *report.CukeElement
	FailedStep *report.CukeStep
}
```

This context is updated for each failed scenario. Its properties are defined as follows:

| Property   | Description                                                                |
| ---------- | -------------------------------------------------------------------------- |
| Feature    | Current Jira feature. The current, failed element is part of this feature. |
| Element    | Current Jira scenario, scenario outline, or background that failed.        |
| FailedStep | Failed step owned by the current element.                                  |

### kosher Cucumber API

The properties of the IssueContext (above) are types defind in kosher's Cucumber API. The full API is defined at [cucumber_format.go](https://github.com/cbush06/kosher/blob/master/report/cucumber_format.go).

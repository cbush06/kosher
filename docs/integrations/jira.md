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
* You have executed the kosher test(s) with a report format that produces a `results.json` file in the `/results` directory (e.g. cucumber, simple, bootstrap, html)
* You have added Jira integration configurations to the `settings.json` file in the `/config` directory

## Command-line Interface

```bash
$ kosher send jira
```

For more details on the `send jira` command, see the [CLI](../../cli.html#jira) page.

### Credentials

After executing the `kosher send jira` command, kosher will prompt you for your username and password. If successful, it will display your full name.

```
$ kosher send jira
Enter Jira credentials...
Username: jdoe
Password: 
Successfully connected to Jira as [Doe, John H]
```

### Selecting a Project

Once you have successfully authenticated, kosher will list all Jira projects you can access and ask you to select one. Your new issues will be created in the project you choose.

```
                JIRA PROJECTS
_____________________________________________
[1]	Kosher
[2]	Golang Practice
[3]	Kosher Docs
_____________________________________________

Select Project: 1
```

### Selecting an Issue Type

After selecting a project, you will be prompted to choose the type of issue to create for failed kosher tests.

```
             JIRA ISSUE TYPES
_____________________________________________
[1]	Bug
[2]	Project Action Item
[3]	Task
[4]	Sub-task
[5]	Epic
[6]	Test Plan
[7]	Test
[8]	Technical task
[9]	Test Case
[10]	Story
[11]	Enhancement
[12]	Project Risk
[13]	Project Opportunity
[14]	Project Issue
[15]	Test Case Template
_____________________________________________

Select Issue Type: 1
```

### Enter an "Affects Version"

Next, kosher prompts you for the "Affects Version." Normally, this is the version of the software project that the bug/issue was discovered in. kosher will either accept blank (nothing) or a valid, existing version that has been created in Jira.

```
Enter "Affects Version": 1.4.0
```

### Enter "Labels"

After the Affects Version, kosher will prompt you to enter "Labels." This should be a comma-delimited list of labels you want added to each of the issues kosher creates in Jira.

```
Enter "Labels": FUNC_TESTING,UX_VALIDATION,UI_VALIDATION,KOSHER
```

### Create Issue for Each Failure

Following the "Affects Version," kosher will prompt you for each failed scenario. If you wish to create an issue for a failure, enter "y", "Y", "yes", or "Yes".

If you choose to create an issue for a failure, kosher will prompt you to select a Priority level for the issue.

If the issue is created successfully, kosher will print the ticket's new key for your reference. It will then repeat this process for the next failure.

```
Create [Verify Results of Navigation Actions: Verify Redirect] (Y/n): y
	Choose priority...
		[1] Blocker
		[2] Critical
		[3] Major
		[4] Normal
		[5] Minor
		[6] Trivial
	Enter priority selection: 2
	Issue [KOSHE-1595] created...

Create [Verify Results of Navigation Actions: Verify Redirect] (Y/n):
```

## Configuration

The Jira integration requires, at a minimum, that the host be specified. Other optional settings may, also, be specified.

### Format

**config/settings.json**

```json
{
    ...,
    "integrations": {
        "jira": {
            "host": "https://jira.myserver.com",
            "summaryTemplate": "jira_summary.txt",
            "descriptionTemplate": "jira_description.txt",
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

### Settings

| Setting             | Description                                                                                                                     |
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| host                | URL of the Jira server to connect to.                                                                                           |
| summaryTemplate     | Name of file in the `/config` directory that contains a Golang template file to be used for creating issue summaries.           |
| descriptionTemplate | Name of file in the `/config` directory that contains a Golang template file to be used for creating issue descriptions.        |
| defaults            | Default values to use rather than prompting the user. These apply if the `--default` flag is used: `kosher send jira --default` |

### Defaults

Providing the `defaults` section allows you to avoid the many prompts given by kosher when sending your results to Jira. However, these are only used if you specify the `--default` flag, as in `kosher send jira --default`.

| Setting        | Default Value if Not Specified | Description                                                                                                      |
| -------------- | ------------------------------ | ---------------------------------------------------------------------------------------------------------------- |
| projectKey     | PROJE                          | Jira project key of the project kosher should create issues in.                                                  |
| issueType      | Bug                            | Issue Type to set for new issues created in Jira.                                                                |
| affectsVersion | 1.0.0                          | Affects version to set for new issues created in Jira. *Set this to blank in the `settings.json` file for none.* |
| labels         | test,functional,kosher         | Labels to add to new issues created in Jira. *Set this to blank in the `settings.json` file for none.*           |
| priority       | Normal                         | Priority to set for new issues created in Jira.                                                                  |

## Templates

kosher's Jira integration creates the summary and description using Golang templates. kosher has default templates embedded in its source, but you may create your own templates in files saved in the `/config` directory. If you create your own template files, they must be specified in the `settings.json` file. See [Configuration Settings](#settings) above for details on specifying template files.

For more information on Golang's templates, see [text/template](https://golang.org/pkg/text/template/).

### Default Templates

**Default summary template**

```
{% raw %}{{.Feature.Name}}: {{.Element.Name}}{% endraw %}
```

The template above results in this text for one of kosher's own tests:

```
Verify Results of Navigation Actions: Verify Redirect
```

**Default description template**

```
{% raw %}h2. Issue:
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
{{end}}{% endraw %}
```

The template above results in this text for one of kosher's own tests (it will appear styled in Jira):

```
Issue:
expected URL to be https://www.seleniumeasy.com/test/basic-checkbox-demo.html but was https://www.seleniumeasy.com/test/basic-first-form-demo.html

Feature Title:
Verify Results of Navigation Actions

Scenario Title:
Verify Redirect

Scenario Description
After clicking on a link, verify redirection to the appropriate page.

Failed Step / Actual Result:
Then I should be redirected to the "basic-checkbox" page
expected URL to be https://www.seleniumeasy.com/test/basic-checkbox-demo.html but was https://www.seleniumeasy.com/test/basic-first-form-demo.html

Steps / Expected Result:
Given I go to the "table-search" page
And I maximize the window
Given I click the "Input Forms" link
And I follow "Simple Form Demo"
Then I should be redirected to the "basic-checkbox" page
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

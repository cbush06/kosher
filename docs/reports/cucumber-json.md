---
layout: default
title: cucumber-json Report Format
description: Description of the "cucumber-json" report format.
parent: Reports
nav_order: 2
---

# cucumber-json Report Format

The `cucumber-json` report format generates a Cucumber-style JSON report named `cucumber.json`.
{: .fs-6 .fw-300 }

## Format

For the `*.feature` file:
```gherkin
Feature: Verify Results of Navigation Actions

    After clicking a link, clicking a button, or submitting
    a form, verfiy that the correct navigations take place.

    Background: Start on the Table Filter page
        Given I go to the "table-search" page
        And I maximize the window

    Scenario: Verify Page
        After clicking on a link, verify that appropriate page is shown.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be on the "bootstrap-date-picker" page

    Scenario: Verify Redirect
        After clicking on a link, verify redirection to the appropriate page.

        Given I click the "Date pickers" link
        And I click the "Bootstrap Date Picker" link
        Then I should be redirected to the "bootstrap-date-picker" page

    Scenario: Verify Javascript Alerts
        Verify that alerts can be verified.

        Given I am on the "js-popup" page
        And I should not see the popup "I am an alert box!"
        When I click the first instance of "Click me!"
        And I wait 1 seconds
        Then I should see the popup "I am an alert box!"
```

The `cucumber-json` report format would output a `cucumber.json` file similar to this:
```json
[
    {
        "uri": "features/navigation_verification.feature",
        "id": "verify-results-of-navigation-actions",
        "keyword": "Feature",
        "name": "Verify Results of Navigation Actions",
        "description": "    After clicking a link, clicking a button, or submitting\n    a form, verfiy that the correct navigations take place.",
        "line": 1,
        "elements": [
            {
                "id": "verify-results-of-navigation-actions;verify-page",
                "keyword": "Scenario",
                "name": "Verify Page",
                "description": "        After clicking on a link, verify that appropriate page is shown.",
                "line": 10,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I go to the \"table-search\" page",
                        "line": 7,
                        "match": {
                            "location": "navigation_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 2293101346
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I maximize the window",
                        "line": 8,
                        "match": {
                            "location": "browser_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 114586972
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "I click the \"Date pickers\" link",
                        "line": 13,
                        "match": {
                            "location": "form_steps.go:322"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 231599608
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I click the \"Bootstrap Date Picker\" link",
                        "line": 14,
                        "match": {
                            "location": "form_steps.go:322"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 954948851
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should be on the \"bootstrap-date-picker\" page",
                        "line": 15,
                        "match": {
                            "location": "navigation_verification_steps.go:15"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 16532090
                        }
                    }
                ]
            },
            {
                "id": "verify-results-of-navigation-actions;verify-redirect",
                "keyword": "Scenario",
                "name": "Verify Redirect",
                "description": "        After clicking on a link, verify redirection to the appropriate page.",
                "line": 17,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I go to the \"table-search\" page",
                        "line": 7,
                        "match": {
                            "location": "navigation_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 128528821
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I maximize the window",
                        "line": 8,
                        "match": {
                            "location": "browser_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 118519611
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "I click the \"Date pickers\" link",
                        "line": 20,
                        "match": {
                            "location": "form_steps.go:322"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 181740551
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I click the \"Bootstrap Date Picker\" link",
                        "line": 21,
                        "match": {
                            "location": "form_steps.go:322"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 407340911
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should be redirected to the \"bootstrap-date-picker\" page",
                        "line": 22,
                        "match": {
                            "location": "navigation_verification_steps.go:15"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 24284158
                        }
                    }
                ]
            },
            {
                "id": "verify-results-of-navigation-actions;verify-javascript-alerts",
                "keyword": "Scenario",
                "name": "Verify Javascript Alerts",
                "description": "        Verify that alerts can be verified.",
                "line": 24,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I go to the \"table-search\" page",
                        "line": 7,
                        "match": {
                            "location": "navigation_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 138181406
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I maximize the window",
                        "line": 8,
                        "match": {
                            "location": "browser_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 107607272
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "I am on the \"js-popup\" page",
                        "line": 27,
                        "match": {
                            "location": "navigation_steps.go:11"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 877850499
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I should not see the popup \"I am an alert box!\"",
                        "line": 28,
                        "match": {
                            "location": "navigation_verification_steps.go:48"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 10514316
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I click the first instance of \"Click me!\"",
                        "line": 29,
                        "match": {
                            "location": "form_steps.go:328"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 161125452
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I wait 1 seconds",
                        "line": 30,
                        "match": {
                            "location": "browser_steps.go:20"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 1000984303
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should see the popup \"I am an alert box!\"",
                        "line": 31,
                        "match": {
                            "location": "navigation_verification_steps.go:48"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 6992803
                        }
                    }
                ]
            }
        ]
    }
]
```
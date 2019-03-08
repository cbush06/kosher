---
layout: default
title: Home
nav_order: 1
description: Web driver wrapper written in Golang that empowers non-developers to functionally test web applications using simple Gherkin scripts.
---
# Automated Testing Without Coding
{: .fs-9 }

kosher tests are a _single source of truth_ for your projects: requirements, documentation, and tests all in one
{: .fs-6 .fw-300 }

<p align="center">![kosher](images/kosher_logo.png)</p>

kosher is a tool that merges several technologies together to automate and simplify functional testing. These include:
* Web Drivers ([https://agouti.org/](https://agouti.org/)): automates tests by remote-controlling browsers
* Cucumber/Gherkin ([https://github.com/DATA-DOG/godog](https://github.com/DATA-DOG/godog)): enables tests scripts to be written in human-readable, BDD-compatible scripts
* Virtual File System ([https://github.com/spf13/afero](https://github.com/spf13/afero)): abstracts away differences between operating systems making the application platform-agnostic and easier to test
* Command-line parser ([https://github.com/spf13/cobra](https://github.com/spf13/cobra)): enables powerful, maintainable, and user-friendly CLI control
* Configuration-file parser ([https://github.com/spf13/viper](https://github.com/spf13/viper)): simplifies use of JSON-based configuration files



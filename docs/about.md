---
layout: default
title: About kosher
nav_order: 70
---

# About kosher

kosher is a tool that merges several technologies together to automate and simplify functional testing. These include:
{: .fs-6 .fw-300 }

* Web Drivers ([https://agouti.org/](https://agouti.org/)): automates tests by remote-controlling browsers
* Cucumber/Gherkin ([https://github.com/DATA-DOG/godog](https://github.com/DATA-DOG/godog)): enables tests scripts to be written in human-readable, BDD-compatible scripts
* Virtual File System ([https://github.com/spf13/afero](https://github.com/spf13/afero)): abstracts away differences between operating systems making the application platform-agnostic and easier to test
* Command-line parser ([https://github.com/spf13/cobra](https://github.com/spf13/cobra)): enables powerful, maintainable, and user-friendly CLI control
* Configuration-file parser ([https://github.com/spf13/viper](https://github.com/spf13/viper)): simplifies use of JSON-based configuration files
* Universal Installer ([http://izpack.org/](http://izpack.org/)): allows a single installer to be developed for all platforms (or in kosher's case, Windows and Linux)
* Accessibility Scanner ([https://github.com/dequelabs/axe-core](https://github.com/dequelabs/axe-core)): scans pages for accessibility violations and concerns

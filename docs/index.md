---
layout: default
title: Home
nav_order: 1
description: Web driver wrapper written in Golang that empowers non-developers to functionally test web applications using simple Gherkin scripts.
---
<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.2/css/all.css" integrity="sha384-fnmOCqbTlWIlj8LyTjo7mOUStjsKC4pOpQbqyi7RrhN7udi9RwhKkMHpvLbHG9Sr" crossorigin="anonymous" />
# Automated Testing Without Coding
{: .fs-9 }

kosher tests are a _single source of truth_ for your projects: requirements, documentation, and tests all in one
{: .fs-6 .fw-300 }

[Get started now](#getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 } [View it on GitHub](https://github.com/cbush06/kosher){: .btn .fs-5 .mb-4 .mb-md-0 } 

---

## Getting Started

### Installing kosher

1. Download the installer from the latest published release of kosher: [https://github.com/cbush06/kosher/releases/](https://github.com/cbush06/kosher/releases/)
2. Run the installer with *administrative privileges*. You can do one of the following:
   * Simply double-click the `kosher.jar` file
   * If double-clicking fails on a linux system, open a terminal and run
```bash
$ sudo java -jar kosher.jar
```
> The installer will suggest several tools to aid in writing Gherkin scripts and running kosher tests. At a minimum, we recommend Chrome, Git (and Git Bash, if you're using Windows), and Visual Studio Code.
3. The installer should automatically add kosher to the PATH. To verify the installation, open a new terminal and key in the command below. Your output should look similar.
```bash
$ kosher version
1.0.0
```

### Trying it Out

1. Create a new directory to house your first kosher project. Using a bash terminal, enter:
```bash
$ mkdir kosher-project
$ cd kosher-project
```
2. Initialize the project (create its basic structure, configuration files, and an example test script):
```bash
$ kosher init KosherProject
Project [KosherProject] initialized...
```
3. Assuming you have Chrome installed, run the example test with the following command:
```bash
$ kosher run features/
```

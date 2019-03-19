---
layout: default
title: Home
nav_order: 0
description: Functionally test your web application through simple, readable Gherkin scripts. No programming required!
---

# Automated Testing Without Coding

Functionally test your web application through simple, readable Gherkin scripts. No programming required!
{: .fs-6 .fw-300 }

[Get started now](#getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 } [View it on GitHub](https://github.com/cbush06/kosher){: .btn .fs-5 .mb-4 .mb-md-0 } 

---

## Getting Started

### Installing kosher

1. Download the installer from the latest published release of kosher: [https://github.com/cbush06/kosher/releases/latest](https://github.com/cbush06/kosher/releases/latest)
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

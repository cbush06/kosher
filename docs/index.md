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
3. The installer should automatically add kosher to the PATH. To verify the installation, open a new terminal and key in the commnad below. Your output should look similar.
```bash
$ kosher version
1.0.0
```
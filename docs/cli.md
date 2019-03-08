---
layout: default
title: CLI
nav_order: 2
description: Documentation on CLI commands for kosher.
---

# CLI
{: .no_toc }

kosher is a command-line utility; thus, you interface with it via the following terminal commands:
{: .fs-6 .fw-300 }

## Commands
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## init

`init` creates the necessary project structure with simple example tests and config files to quickly get you started.

### Format

```bash
$ kosher init [flags] [project name] [directory]
```

### Flags

| Flag           | Description                                                                                                         |
| -------------- | ------------------------------------------------------------------------------------------------------------------- |
| -e, --empty    | init creates the necessary project structure with simple example tests and config files to quickly get you started. |
| -f, --force    | Create a project inside a non-empty directory.                                                                      |
| -h, --help     | help for init                                                                                                       |
| -p, --platform | Set the platform: desktop, web (default "web")                                                                      |


### Arguments

| Argument     | Required | Description                                                     |
| ------------ | -------- | --------------------------------------------------------------- |
| project name | yes      | Name of the new project being initialized.                      |
| directory    | no       | Relative path to subdirectory to initialize the new project in. |

## run

`run` executes one or more `*.feature` files of the current project. **This command must be executed from the root of the kosher project.**

### Format

```bash
$ kosher run [flags] [path]
```

### Flags

| Flag              | Description                                                                                                              |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------ |
| --appVersion      | Version of the application being tested. This will be used when creating the report after testing has completed.         |
| -e, --environment | Sets the environment to execute tests against. This overrides the default environemnt specified in config/settings.json. |
| -h, --help        | help for run                                                                                                             |

### Arguments

| Argument | Required | Description                                                                                                                                                                 |
| -------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| path     | no       | Providing this argument restricts the selection of `*.feature` files executed to a specific subfolder or even a specific file. Standard Linux _glob_ patterns are accepted. |

## version

`version` prints the version of kosher running to the terminal.

### Format

```bash
$ kosher version [flags]
```

### Flags

| Flag       | Description      |
| ---------- | ---------------- |
| -h, --help | help for version |

### Arguments

> There are no arguments for `version`.

## help

`help` prints the usage instructions for kosher to the terminal.

### Format

```bash
$ kosher help [flags] [command]
```

### Flags

| Flag       | Description   |
| ---------- | ------------- |
| -h, --help | help for help |

### Arguments

| Argument | Required | Description                                       |
| -------- | -------- | ------------------------------------------------- |
| command  | no       | Specifies a kosher command to print the help for. |
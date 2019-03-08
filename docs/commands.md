---
layout: default
title: CLI
nav_order: 2
---

# CLI
{: .no_toc }

kosher is a command-line utility; thus, you interface with it via the following terminal commands:
{: .fs-6 .fw-300 }

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
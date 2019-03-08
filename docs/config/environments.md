---
layout: default
title: environments.json
description: Format of environments.json
parent: Configuration
---

# environments.json

The `environments.json` file allows the user to specify different testing environments along with the base URLs of those environments. The [settings.json](settings.html) file can specify a `defaultEnvironment`. The `run` command's `--environment` flag can specify an environment to use for that test run.
{: .fs-6 .fw-300 }

## Format

```json
{
    "production": "http://www.your-production-env.com/",
    "test": "https://www.seleniumeasy.com/test",
    "dev": "http://www.your-dev-env.com",
    "uat": "http://www.your-uat-env.com"
}
```
---
layout: default
title: Command Line installation (Admin Access MANDATORY) 
nav_order: 43
---

## Getting Started

### Installing kosher

1. Download the installer from the latest published release of kosher: [https://github.com/cbush06/kosher/releases/latest](https://github.com/cbush06/kosher/releases/latest)
2. Run the installer with *administrative privileges*. You can do one of the following:
   * Simply double-click the `kosher.jar` file (*this may work for some*)
   * If double-clicking fails on a linux system, open a terminal and run `sudo java -jar kosher.jar`
   * If double-clicking fails on a Windows system, skip to the next section **Install using Windows Command**

### Install using Windows Command Line

1. Navigate to the Windows Command Line shortcut, right-click, and Open as Administrator

    ![Command Line]({{site.baseurl}}/assets/images/navigatingtoCMD.png)

2. When the terminal opens Chage Directory (cd) into the following file path: (*C:\Users\user\Downloads*) 

    a. ![Command Line]({{site.baseurl}}/assets/images/CD1.png)
    
    b. ![Command Line]({{site.baseurl}}/assets/images/CD2.png)

3. Execute the following command to begin the install: (*java -jar "kosher(version#).jar*)

    ![Command Line]({{site.baseurl}}/assets/images/javainstall.png)

4. The kosher installation wizard should appear. Follow its prompts to complete installation.

    ![Command Line]({{site.baseurl}}/assets/images/installer.png)

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
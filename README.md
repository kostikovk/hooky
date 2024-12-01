# Hooky CLI [WIP]

**Hooky CLI** is a command-line tool for managing Git hooks in Go projects. Designed with simplicity and efficiency in mind, Hooky makes it easy to set up, configure, and run Git hooks such as pre-commit, pre-push, and more. Whether you’re enforcing coding standards, running tests, or automating tasks, Hooky streamlines the process in a Go-centric development workflow.

## Features
- **Easy Hook Management**: Install and configure Git hooks effortlessly.
- **Go-Friendly**: Optimized for Go projects and workflows.
- **Customizable Hooks**: Define and execute your custom Git hooks.
- **Seamless Integration**: Works out of the box with Git repositories.

## Installation
To install Hooky CLI, run the following command:

```bash
go install github.com/kostikovk/hooky@latest
```
**Note:** Replace @latest with a specific version if you want to install a particular release.

## Usage
To get started with Hooky CLI, run the following command:

```bash
hooky init
```
This command initializes Hooky in your Git repository and creates a `.hooky` directory with default hooks.

## Install specific hook
To install a specific hook, run the following command:

```bash
hooky install pre-commit
```
This command installs the `pre-commit` hook in your Git repository.
If it's already installed, it will overwrite the existing hook.

# Hooky CLI

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
If a conflicting hook already exists in `.git/hooks`, Hooky stops and reports it (non-destructive by default).
Use `--force` to replace conflicting hooks and `--backup` (enabled by default) to keep backups:

```bash
hooky install pre-commit --force --backup
```

## Husky-like add command
To add or update a hook with a custom command:

```bash
hooky add pre-commit "go test ./..."
```

This writes the hook script to `.hooky/git-hooks/pre-commit` and syncs it into `.git/hooks`.
If `.git/hooks/pre-commit` already exists and conflicts, use:

```bash
hooky add pre-commit "go test ./..." --force --backup
```

## Doctor command
To verify Hooky installation and hook wiring:

```bash
hooky doctor
```

`doctor` checks:
- `.git` exists
- `.hooky` and `.hooky/git-hooks` exist
- each Hooky hook is installed in `.git/hooks` as a symlink to `.hooky/git-hooks`

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear description of your changes.

## License
This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE.md) file for details.

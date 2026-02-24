# Hooky — Husky-style Git hooks for Go

**Hooky CLI** is a command-line tool for managing Git hooks in Go projects. Designed with simplicity and efficiency in mind, Hooky makes it easy to set up, configure, and run Git hooks such as pre-commit, pre-push, and more. Whether you’re enforcing coding standards, running tests, or automating tasks, Hooky streamlines the process in a Go-centric development workflow.

## Features
- **Easy Hook Management**: Install and configure Git hooks effortlessly.
- **Go-Friendly**: Optimized for Go projects and workflows.
- **Customizable Hooks**: Define and execute your custom Git hooks.
- **Seamless Integration**: Works out of the box with Git repositories.
- **Safe by Default**: Non-destructive sync behavior for `.git/hooks`.

## Installation
To install Hooky CLI, run the following command:

```bash
go install github.com/kostikovk/hooky@latest
```
**Note:** Replace @latest with a specific version if you want to install a particular release.

## Quick Start
Initialize Hooky in your repository:

```bash
hooky init
```
This creates `.hooky/hooks`, adds a default `pre-commit` hook, and syncs hooks into `.git/hooks`.

Add or update your own hook command (Husky-style):

```bash
hooky add pre-commit "go test ./..."
```

Check setup health:

```bash
hooky doctor
```

## Command Reference

### `hooky init`
Initializes Hooky for the current repository and syncs hook links into `.git/hooks`.

```bash
hooky init [--force] [--backup]
```

### `hooky install [hook]`
Creates a predefined hook in `.hooky/hooks` and syncs it to `.git/hooks`.

```bash
hooky install pre-commit
```

```bash
hooky install pre-commit --force --backup
```

### `hooky add [hook] [command]`
Creates or updates a hook script with a custom command, then syncs it.

```bash
hooky add pre-commit "go test ./..."
```

```bash
hooky add pre-commit "go test ./..." --force --backup
```

### `hooky doctor`
Verifies Hooky installation and `.git/hooks` wiring.

```bash
hooky doctor
```

Doctor checks:
- `.git` exists
- `.hooky` and `.hooky/hooks` exist
- each Hooky hook is installed in `.git/hooks` as a symlink to `.hooky/hooks`

### `hooky list`
Lists available hooks.

```bash
hooky list
hooky list --installed
```

### Other commands
```bash
hooky uninstall
hooky version
hooky --help
```

## Safety Model
Hooky sync is **non-destructive by default**:
- If a conflicting file/symlink already exists in `.git/hooks`, Hooky reports conflict and exits.
- Use `--force` to replace conflicting entries.
- Use `--backup` (default: `true` with `--force`) to preserve replaced hooks as `*.hooky.bak`.

## Recovery Notes
- If `.git/hooks` is deleted manually, your source hooks in `.hooky/hooks` remain intact.
- Re-run one of the following to recreate links:
  - `hooky init`
  - `hooky install <hook>`
  - `hooky add <hook> "<command>"`
- Run `hooky doctor` to confirm health.

## Automated Releases (GitHub Actions)
This repository uses GitHub Actions for CI and automated semantic versioning:
- `CI`: runs tests on pull requests and pushes to `main`.
- `Release Please`: on `main`, computes stable versions and creates GitHub releases/tags.
- `Release Assets`: builds binaries for Linux, macOS, and Windows and uploads them to each GitHub release.

Version bumps follow Conventional Commits:
- `fix: ...` -> patch bump (`v1.2.3` -> `v1.2.4`)
- `feat: ...` -> minor bump (`v1.2.3` -> `v1.3.0`)
- `feat!: ...` or `BREAKING CHANGE: ...` -> major bump (`v1.2.3` -> `v2.0.0`)

Use a semantic PR title (or squash-merge commit title), for example:
- `fix: handle missing .git/hooks directory`
- `feat: add doctor command symlink validation`
- `feat!: change init defaults`

Note: to ensure release-created events can trigger other workflows (asset upload), set a repository secret named `RELEASE_PLEASE_TOKEN` with a PAT that has `contents` and `pull_request` permissions.

## Development
- Go version: `1.26.0` (see `go.mod`)
- Lint config: `.golangci.yml`
- Versioning: build-time injected via linker flags (`cmd.Version`)

Common tasks:
```bash
go test ./...
golangci-lint run -c .golangci.yml
make test
make build
make build VERSION=v1.0.2
```

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear description of your changes.

## License
This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE.md) file for details.

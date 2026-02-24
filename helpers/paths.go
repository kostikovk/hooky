package helpers

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Test overrides. When empty, paths are resolved dynamically from the git repository.
var AbsoluteHookyPath string
var AbsoluteHookyGitHooksPath string
var AbsoluteGitPath string
var AbsoluteGitHooksPath string

func gitTopLevel() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to detect git root: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

func gitDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--git-dir")

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to detect git dir: %w", err)
	}

	dir := strings.TrimSpace(string(out))
	if filepath.IsAbs(dir) {
		return dir, nil
	}

	root, err := gitTopLevel()
	if err != nil {
		return "", err
	}

	return filepath.Clean(filepath.Join(root, dir)), nil
}

func getHookyPath() (string, error) {
	if AbsoluteHookyPath != "" {
		return AbsoluteHookyPath, nil
	}

	root, err := gitTopLevel()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, ".hooky"), nil
}

func getHookyGitHooksPath() (string, error) {
	if AbsoluteHookyGitHooksPath != "" {
		return AbsoluteHookyGitHooksPath, nil
	}

	hookyPath, err := getHookyPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(hookyPath, "git-hooks"), nil
}

func getGitPath() (string, error) {
	if AbsoluteGitPath != "" {
		return AbsoluteGitPath, nil
	}

	return gitDir()
}

func getGitHooksPath() (string, error) {
	if AbsoluteGitHooksPath != "" {
		return AbsoluteGitHooksPath, nil
	}

	gitPath, err := getGitPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(gitPath, "hooks"), nil
}

func HookyGitHooksPath() (string, error) {
	return getHookyGitHooksPath()
}

func GitHooksPath() (string, error) {
	return getGitHooksPath()
}

package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// GitHooks List of Git hooks.
// Source: https://git-scm.com/docs/githooks
var GitHooks = []string{
	"applypatch-msg",
	"commit-msg",
	"post-applypatch",
	"post-checkout",
	"post-commit",
	"post-merge",
	"post-receive",
	"post-rewrite",
	"post-update",
	"pre-applypatch",
	"pre-auto-gc",
	"pre-commit",
	"pre-push",
	"pre-rebase",
	"pre-receive",
	"prepare-commit-msg",
	"push-to-checkout",
	"update",
}

var AbsoluteGitPath = getAbsolutePath(".git")
var AbsoluteGitHooksPath = getAbsolutePath(".git/hooks")

// GitHookExists checks if a Git hook exists in the GitHooks slice.
func GitHookExists(hook string) bool {
	return Contains(GitHooks, func(h string) bool {
		return h == hook
	})
}

// IsGitRepository checks if the current directory is a Git repository.
func IsGitRepository() bool {
	return dirExists(AbsoluteGitPath)
}

// InitGit initializes a Git repository.
func InitGit() error {
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to initialize Git repository: %w", err)
	}

	return nil
}

// PromptToInitGit prompts the user to initialize a Git repository.
func PromptToInitGit() error {
	return Prompt("This is not a Git repository. Would you like to initialize it?")
}

// PromptToCopyGitHooksToHooky prompts the user to copy Git hooks to Hooky repository.
func PromptToCopyGitHooksToHooky() error {
	return Prompt("Would you like to copy Git hooks to Hooky repository?")
}

// DeleteGitHooksDirectory .git/hooks folder with all its contents.
func DeleteGitHooksDirectory() error {
	return os.RemoveAll(AbsoluteGitHooksPath)
}

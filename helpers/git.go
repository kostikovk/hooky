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

func GitHookExists(hook string) bool {
	return Contains(GitHooks, func(h string) bool {
		return h == hook
	})
}

func IsGitRepository() bool {
	return dirExists(AbsoluteGitPath)
}

func InitGit() error {
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize Git repository: %w", err)
	}

	return nil
}

func PromptToInitGit() error {
	return Prompt("This is not a Git repository. Would you like to initialize it?")
}

func PromptToCopyGitHooksToHooky() error {
	return Prompt("Would you like to copy Git hooks to Hooky repository?")
}

func DeleteGitHooksDirectory() error {
	return os.RemoveAll(AbsoluteGitHooksPath)
}

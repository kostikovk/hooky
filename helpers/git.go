package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// List of Git hooks.
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

// GitHookExists checks if a Git hook exists.
func GitHookExists(hook string) bool {
	for _, h := range GitHooks {
		if h == hook {
			return true
		}
	}

	return false
}

// IsGitRepository checks if the current directory is a Git repository.
func IsGitRepository() bool {
	return dirExists(AbsoluteGitPath)
}

// HasGitHooksDirectory checks if the current directory has a .git/hooks folder.
func HasGitHooksDirectory() bool {
	return dirExists(AbsoluteGitHooksPath)
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
	fmt.Println("This is not a Git repository. Would you like to initialize it?")
	fmt.Print("Y/n: ")

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return fmt.Errorf("Failed to read response: %w", err)
	}

	if response == "Y" || response == "y" {
		return InitGit()
	}

	return nil
}

// Delete .git/hooks folder with all its contents.
func DeleteGitHooksDirectory() error {
	return os.RemoveAll(AbsoluteGitHooksPath)
}

// Check if .git/hooks has some hooks.
func HasGitHooks() bool {
	files, err := os.ReadDir(AbsoluteGitHooksPath)
	if err != nil {
		return false
	}

	for _, file := range files {
		if GitHookExists(file.Name()) {
			return true
		}
	}

	return false
}

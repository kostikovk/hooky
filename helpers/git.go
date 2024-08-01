package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// PromptToCopyGitHooksToGoHooks prompts the user to copy Git hooks to GoHooks repository.
func PromptToCopyGitHooksToGoHooks() error {
	fmt.Println("Would you like to copy Git hooks to GoHooks repository?")
	fmt.Print("Y/n: ")

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return fmt.Errorf("Failed to read response: %w", err)
	}

	if response == "Y" || response == "y" {
		return CopyGitHooksToGoHooks()
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

	hasGitHook := false
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sample") {
			hasGitHook = true
			break
		}
	}

	return hasGitHook
}

// CopyGitHooksToGoHooks copies Git hooks to GoHooks repository.
func CopyGitHooksToGoHooks() error {
	files, err := os.ReadDir(AbsoluteGitHooksPath)
	if err != nil {
		return fmt.Errorf("Failed to read .git/hooks directory: %w", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sample") {
			continue
		}

		pathFrom := fmt.Sprintf("%s/%s", AbsoluteGitHooksPath, file.Name())
		pathTo := fmt.Sprintf("%s/%s", AbsoluteGoHooksGitHooksPath, file.Name())

		err := copyFile(pathFrom, pathTo)
		if err != nil {
			return fmt.Errorf("Failed to copy Git hook %s to GoHooks: %w", file.Name(), err)
		}
	}

	return nil
}

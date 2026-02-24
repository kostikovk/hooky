package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var (
	hasHookyHooksDirectory = helpers.HasGitHooksDirectory
	listHookyHooks         = helpers.ListOfInstalledGitHooks
)

func RunDoctor(cmd *cobra.Command, args []string) error {
	var issues []string

	if !isGitRepository() {
		issues = append(issues, "not a Git repository (.git not found)")
	}

	if !isHookyRepository() {
		issues = append(issues, "Hooky repository not found (.hooky)")
	}

	if isHookyRepository() && !hasHookyHooksDirectory() {
		issues = append(issues, "Hooky hooks directory not found (.hooky/hooks)")
	}

	if len(issues) > 0 {
		printIssues(cmd, issues)
		return fmt.Errorf("doctor found %d issue(s)", len(issues))
	}

	hooks, err := listHookyHooks()
	if err != nil {
		return fmt.Errorf("failed to read hooky hooks: %w", err)
	}

	for _, hook := range hooks {
		if hook == "" {
			continue
		}

		if !helpers.GitHookExists(hook) {
			continue
		}

		ok, issue := isHookInstalledAndManaged(hook)
		if !ok {
			issues = append(issues, issue)
		}
	}

	if len(issues) > 0 {
		printIssues(cmd, issues)
		return fmt.Errorf("doctor found %d issue(s)", len(issues))
	}

	cmd.Println("Doctor check passed: Hooky is healthy.")
	return nil
}

func isHookInstalledAndManaged(hook string) (bool, string) {
	hookyGitHooksPath, err := helpers.HookyGitHooksPath()
	if err != nil {
		return false, fmt.Sprintf("%s: cannot resolve .hooky/hooks path: %v", hook, err)
	}

	gitHooksPath, err := helpers.GitHooksPath()
	if err != nil {
		return false, fmt.Sprintf("%s: cannot resolve .git/hooks path: %v", hook, err)
	}

	source := filepath.Join(hookyGitHooksPath, hook)
	target := filepath.Join(gitHooksPath, hook)

	info, err := os.Lstat(target)
	if os.IsNotExist(err) {
		return false, fmt.Sprintf("%s: missing from .git/hooks", hook)
	}
	if err != nil {
		return false, fmt.Sprintf("%s: cannot inspect target hook: %v", hook, err)
	}

	if info.Mode()&os.ModeSymlink == 0 {
		return false, fmt.Sprintf("%s: existing hook is not a symlink managed by Hooky", hook)
	}

	link, err := os.Readlink(target)
	if err != nil {
		return false, fmt.Sprintf("%s: cannot read symlink: %v", hook, err)
	}

	if !filepath.IsAbs(link) {
		link = filepath.Join(filepath.Dir(target), link)
	}

	if filepath.Clean(link) != filepath.Clean(source) {
		return false, fmt.Sprintf("%s: symlink points outside .hooky/hooks", hook)
	}

	return true, ""
}

func printIssues(cmd *cobra.Command, issues []string) {
	cmd.Println("Doctor found issues:")
	for _, issue := range issues {
		cmd.Printf("- %s\n", strings.TrimSpace(issue))
	}
}

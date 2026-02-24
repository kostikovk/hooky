package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var AbsoluteHookyPath = getAbsolutePath(".hooky")
var AbsoluteHookyGitHooksPath = getAbsolutePath(".hooky/git-hooks")

type InstallOptions struct {
	Force  bool
	Backup bool
}

func IsHookyRepository() bool {
	return dirExists(AbsoluteHookyPath)
}

func HasGitHooksDirectory() bool {
	return dirExists(AbsoluteHookyGitHooksPath)
}

func CreateHookyGitDirectory() error {
	return os.MkdirAll(AbsoluteHookyGitHooksPath, 0750)
}

func DeleteHookyDirectory() error {
	return os.RemoveAll(AbsoluteHookyPath)
}

func CreateGitHook(hook string, cmd string) error {
	return writeGitHook(hook, cmd, false)
}

func UpsertGitHook(hook string, cmd string) error {
	return writeGitHook(hook, cmd, true)
}

func writeGitHook(hook string, cmd string, allowOverwrite bool) error {
	if !IsHookyRepository() {
		fmt.Println("Hooky repository not found")
		fmt.Println("Please, do 'hooky init' to create a Hooky repository")

		return fmt.Errorf("hooky repository not found")
	}

	if !HasGitHooksDirectory() {
		fmt.Println("Git hooks directory not found in Hooky repository '.hooky/git-hooks'")
		fmt.Println("Please, do 'hooky uninstall' and 'hooky init' to create a Hooky repository again")

		return fmt.Errorf("git hooks directory not found in Hooky repository '.hooky/git-hooks'")
	}

	if !GitHookExists(hook) {
		return fmt.Errorf("invalid Git hook: %s", hook)
	}

	target := filepath.Join(AbsoluteHookyGitHooksPath, hook)
	if !allowOverwrite && exists(target) {
		return fmt.Errorf("hook already exists: %s", hook)
	}

	content := "#!/bin/sh\n" + cmd + "\n"
	if err := os.WriteFile(target, []byte(content), 0750); err != nil {
		return fmt.Errorf("failed to write hook file: %w", err)
	}

	if err := os.Chmod(target, 0750); err != nil {
		return fmt.Errorf("failed to change file permissions: %w", err)
	}

	return nil
}

func InstallHooks(options InstallOptions) error {
	if !IsHookyRepository() {
		return fmt.Errorf("GoHooks repository not found")
	}

	hooks, err := os.ReadDir(AbsoluteHookyGitHooksPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	err = os.MkdirAll(AbsoluteGitHooksPath, 0750)
	if err != nil {
		return fmt.Errorf("failed to create Git hooks directory: %w", err)
	}

	var conflicts []string
	for _, hook := range hooks {
		if hook.IsDir() || !GitHookExists(hook.Name()) {
			continue
		}

		source := filepath.Join(AbsoluteHookyGitHooksPath, hook.Name())
		target := filepath.Join(AbsoluteGitHooksPath, hook.Name())

		if shouldReplace, err := shouldReplaceHookTarget(target, source); err != nil {
			return err
		} else if !shouldReplace {
			continue
		}

		if exists(target) {
			if !options.Force {
				conflicts = append(conflicts, hook.Name())
				continue
			}

			if options.Backup {
				if _, err := backupHook(target); err != nil {
					return err
				}
			} else if err := os.Remove(target); err != nil {
				return fmt.Errorf("failed to remove existing hook %q: %w", hook.Name(), err)
			}
		}

		if err := os.Symlink(source, target); err != nil {
			return fmt.Errorf("failed to link file: %w", err)
		}

		if err := os.Chmod(target, 0750); err != nil {
			return fmt.Errorf("failed to change file permissions: %w", err)
		}
	}

	if len(conflicts) > 0 {
		return fmt.Errorf(
			"existing git hooks detected: %s (re-run with --force to replace, optionally --backup)",
			strings.Join(conflicts, ", "),
		)
	}

	return nil
}

func shouldReplaceHookTarget(target, source string) (bool, error) {
	info, err := os.Lstat(target)
	if os.IsNotExist(err) {
		return true, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to stat hook target %q: %w", target, err)
	}

	if info.Mode()&os.ModeSymlink == 0 {
		return true, nil
	}

	link, err := os.Readlink(target)
	if err != nil {
		return false, fmt.Errorf("failed to inspect hook symlink %q: %w", target, err)
	}

	if !filepath.IsAbs(link) {
		link = filepath.Join(filepath.Dir(target), link)
	}

	resolvedLink := filepath.Clean(link)
	resolvedSource := filepath.Clean(source)
	return resolvedLink != resolvedSource, nil
}

func backupHook(target string) (string, error) {
	backupPath := target + ".hooky.bak"
	if exists(backupPath) {
		for i := 1; ; i++ {
			candidate := fmt.Sprintf("%s.%d", backupPath, i)
			if !exists(candidate) {
				backupPath = candidate
				break
			}
		}
	}

	if err := os.Rename(target, backupPath); err != nil {
		return "", fmt.Errorf("failed to backup existing hook %q: %w", target, err)
	}
	return backupPath, nil
}

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

func ListOfInstalledGitHooks() ([]string, error) {
	files, err := os.ReadDir(AbsoluteHookyGitHooksPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var hooks []string
	for _, file := range files {
		hooks = append(hooks, file.Name())
	}

	return hooks, nil
}

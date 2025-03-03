package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

var AbsoluteHookyPath = getAbsolutePath(".hooky")
var AbsoluteHookyGitHooksPath = getAbsolutePath(".hooky/git-hooks")

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

	files, err := os.ReadDir(AbsoluteHookyGitHooksPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	if ContainsFile(files, hook) {
		return fmt.Errorf("hook already exists: %s", hook)
	}

	file, err := os.Create(filepath.Join(AbsoluteHookyGitHooksPath, hook))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if _, err = file.WriteString("#!/bin/sh\n" + cmd); err != nil {
		return err
	}

	return nil
}

func InstallHooks() error {
	if !IsHookyRepository() {
		return fmt.Errorf("GoHooks repository not found")
	}

	hooks, err := os.ReadDir(AbsoluteHookyGitHooksPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	err = DeleteGitHooksDirectory()
	if err != nil {
		return fmt.Errorf("failed to delete Git hooks directory: %w", err)
	}

	err = os.MkdirAll(AbsoluteGitHooksPath, 0750)
	if err != nil {
		return fmt.Errorf("failed to create Git hooks directory: %w", err)
	}

	for _, hook := range hooks {
		if hook.IsDir() || !GitHookExists(hook.Name()) {
			continue
		}

		err = os.Symlink(
			filepath.Join(AbsoluteHookyGitHooksPath, hook.Name()),
			filepath.Join(AbsoluteGitHooksPath, hook.Name()),
		)
		if err != nil {
			return fmt.Errorf("failed to link file: %w", err)
		}

		err = os.Chmod(filepath.Join(AbsoluteGitHooksPath, hook.Name()), 0750)
		if err != nil {
			return fmt.Errorf("failed to change file permissions: %w", err)
		}
	}

	return nil
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

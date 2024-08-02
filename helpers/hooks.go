package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

var AbsoluteGoHooksPath = getAbsolutePath(".gohooks")
var AbsoluteGoHooksGitHooksPath = getAbsolutePath(".gohooks/git/hooks")

// IsGoHooksRepository checks if the current directory is a GoHooks repository.
func IsGoHooksRepository() bool {
	return dirExists(AbsoluteGoHooksPath)
}

// CreateGoHooksGitDirectory creates a .gohooks/git/hooks folder.
func CreateGoHooksGitDirectory() error {
	return os.MkdirAll(AbsoluteGoHooksGitHooksPath, 0750)
}

// DeleteGoHooksDirectory .gohooks directory
func DeleteGoHooksDirectory() error {
	return os.RemoveAll(AbsoluteGoHooksPath)
}

// CreateGitHook creates a GoHooks Git hook.
func CreateGitHook(hook string, cmd string) error {
	// check if hook is valid Git hook.
	if !GitHookExists(hook) {
		return fmt.Errorf("invalid Git hook: %s", hook)
	}

	// check if GoHooks Git directory exists.
	files, err := os.ReadDir(AbsoluteGoHooksGitHooksPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// check if hook already exists.
	if ContainsFile(files, hook) {
		return fmt.Errorf("hook already exists: %s", hook)
	}

	// create hook file.
	file, err := os.Create(filepath.Join(AbsoluteGoHooksGitHooksPath, hook))
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

// InstallHooks installs all GoHooks Git hooks.
func InstallHooks() error {
	if !IsGoHooksRepository() {
		return fmt.Errorf("GoHooks repository not found")
	}

	hooks, err := os.ReadDir(AbsoluteGoHooksGitHooksPath)
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

		err = os.Link(
			filepath.Join(AbsoluteGoHooksGitHooksPath, hook.Name()),
			filepath.Join(AbsoluteGitHooksPath, hook.Name()),
		)
		if err != nil {
			return fmt.Errorf("failed to link file: %w", err)
		}

		// make hook executable.
		err = os.Chmod(filepath.Join(AbsoluteGitHooksPath, hook.Name()), 0750)
		if err != nil {
			return fmt.Errorf("failed to change file permissions: %w", err)
		}
	}

	return nil
}

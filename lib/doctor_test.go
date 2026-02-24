package lib

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func TestRunDoctorHealthy(t *testing.T) {
	restore := setupDoctorTestState(t)
	defer restore()

	isGitRepository = func() bool { return true }
	isHookyRepository = func() bool { return true }
	hasHookyHooksDirectory = func() bool { return true }
	listHookyHooks = func() ([]string, error) {
		return []string{"pre-commit"}, nil
	}

	source := filepath.Join(helpers.AbsoluteHookyGitHooksPath, "pre-commit")
	target := filepath.Join(helpers.AbsoluteGitHooksPath, "pre-commit")
	writeFile(t, source, "#!/bin/sh\necho test\n")
	if err := symlink(source, target); err != nil {
		t.Fatalf("create symlink: %v", err)
	}

	cmd := &cobra.Command{}
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)

	if err := RunDoctor(cmd, nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !strings.Contains(out.String(), "Doctor check passed") {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunDoctorDetectsMissingHookInstall(t *testing.T) {
	restore := setupDoctorTestState(t)
	defer restore()

	isGitRepository = func() bool { return true }
	isHookyRepository = func() bool { return true }
	hasHookyHooksDirectory = func() bool { return true }
	listHookyHooks = func() ([]string, error) {
		return []string{"pre-commit"}, nil
	}

	source := filepath.Join(helpers.AbsoluteHookyGitHooksPath, "pre-commit")
	writeFile(t, source, "#!/bin/sh\necho test\n")

	cmd := &cobra.Command{}
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)

	err := RunDoctor(cmd, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(out.String(), "missing from .git/hooks") {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunDoctorIgnoresNonHookFilesInHookyDirectory(t *testing.T) {
	restore := setupDoctorTestState(t)
	defer restore()

	isGitRepository = func() bool { return true }
	isHookyRepository = func() bool { return true }
	hasHookyHooksDirectory = func() bool { return true }
	listHookyHooks = func() ([]string, error) {
		return []string{"pre-commit", "README.md", "tmp-file"}, nil
	}

	source := filepath.Join(helpers.AbsoluteHookyGitHooksPath, "pre-commit")
	target := filepath.Join(helpers.AbsoluteGitHooksPath, "pre-commit")
	writeFile(t, source, "#!/bin/sh\necho test\n")
	if err := symlink(source, target); err != nil {
		t.Fatalf("create symlink: %v", err)
	}

	cmd := &cobra.Command{}
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)

	if err := RunDoctor(cmd, nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !strings.Contains(out.String(), "Doctor check passed") {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunDoctorDetectsRepoIssues(t *testing.T) {
	restore := setupDoctorTestState(t)
	defer restore()

	isGitRepository = func() bool { return false }
	isHookyRepository = func() bool { return false }

	cmd := &cobra.Command{}
	out := bytes.NewBuffer(nil)
	cmd.SetOut(out)

	err := RunDoctor(cmd, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(out.String(), "not a Git repository") {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func setupDoctorTestState(t *testing.T) func() {
	t.Helper()

	origIsGitRepository := isGitRepository
	origIsHookyRepository := isHookyRepository
	origHasHookyHooksDirectory := hasHookyHooksDirectory
	origListHookyHooks := listHookyHooks

	origHookyPath := helpers.AbsoluteHookyPath
	origHookyGitHooksPath := helpers.AbsoluteHookyGitHooksPath
	origGitPath := helpers.AbsoluteGitPath
	origGitHooksPath := helpers.AbsoluteGitHooksPath

	root := t.TempDir()
	helpers.AbsoluteHookyPath = filepath.Join(root, ".hooky")
	helpers.AbsoluteHookyGitHooksPath = filepath.Join(root, ".hooky", "git-hooks")
	helpers.AbsoluteGitPath = filepath.Join(root, ".git")
	helpers.AbsoluteGitHooksPath = filepath.Join(root, ".git", "hooks")

	mkdirAll(t, helpers.AbsoluteHookyGitHooksPath)
	mkdirAll(t, helpers.AbsoluteGitHooksPath)

	return func() {
		isGitRepository = origIsGitRepository
		isHookyRepository = origIsHookyRepository
		hasHookyHooksDirectory = origHasHookyHooksDirectory
		listHookyHooks = origListHookyHooks

		helpers.AbsoluteHookyPath = origHookyPath
		helpers.AbsoluteHookyGitHooksPath = origHookyGitHooksPath
		helpers.AbsoluteGitPath = origGitPath
		helpers.AbsoluteGitHooksPath = origGitHooksPath
	}
}

func mkdirAll(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o750); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o750); err != nil {
		t.Fatalf("write file %s: %v", path, err)
	}
}

func symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

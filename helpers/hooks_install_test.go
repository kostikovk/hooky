package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInstallHooksNonDestructiveByDefault(t *testing.T) {
	root := t.TempDir()
	restorePaths := useTestPaths(t, root)
	defer restorePaths()

	source := filepath.Join(AbsoluteHookyGitHooksPath, "pre-commit")
	target := filepath.Join(AbsoluteGitHooksPath, "pre-commit")
	if err := os.WriteFile(source, []byte("#!/bin/sh\necho hooky\n"), 0o750); err != nil {
		t.Fatalf("write source hook: %v", err)
	}
	if err := os.WriteFile(target, []byte("#!/bin/sh\necho legacy\n"), 0o750); err != nil {
		t.Fatalf("write target hook: %v", err)
	}

	err := InstallHooks(InstallOptions{})
	if err == nil {
		t.Fatal("expected conflict error, got nil")
	}
	if !strings.Contains(err.Error(), "existing git hooks detected") {
		t.Fatalf("unexpected error: %v", err)
	}

	info, err := os.Lstat(target)
	if err != nil {
		t.Fatalf("stat target: %v", err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		t.Fatal("target hook should remain untouched (not symlink) without --force")
	}
	content, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("read target: %v", err)
	}
	if string(content) != "#!/bin/sh\necho legacy\n" {
		t.Fatalf("target hook content changed unexpectedly: %q", string(content))
	}
}

func TestInstallHooksForceWithBackup(t *testing.T) {
	root := t.TempDir()
	restorePaths := useTestPaths(t, root)
	defer restorePaths()

	source := filepath.Join(AbsoluteHookyGitHooksPath, "pre-commit")
	target := filepath.Join(AbsoluteGitHooksPath, "pre-commit")
	if err := os.WriteFile(source, []byte("#!/bin/sh\necho hooky\n"), 0o750); err != nil {
		t.Fatalf("write source hook: %v", err)
	}
	if err := os.WriteFile(target, []byte("#!/bin/sh\necho legacy\n"), 0o750); err != nil {
		t.Fatalf("write target hook: %v", err)
	}

	if err := InstallHooks(InstallOptions{Force: true, Backup: true}); err != nil {
		t.Fatalf("install hooks: %v", err)
	}

	link, err := os.Readlink(target)
	if err != nil {
		t.Fatalf("target should be symlink: %v", err)
	}
	if filepath.Clean(link) != filepath.Clean(source) {
		t.Fatalf("unexpected symlink target: %q, want %q", link, source)
	}

	backupPath := target + ".hooky.bak"
	backupContent, err := os.ReadFile(backupPath)
	if err != nil {
		t.Fatalf("backup not found: %v", err)
	}
	if string(backupContent) != "#!/bin/sh\necho legacy\n" {
		t.Fatalf("unexpected backup content: %q", string(backupContent))
	}
}

func TestInstallHooksForceWithoutBackup(t *testing.T) {
	root := t.TempDir()
	restorePaths := useTestPaths(t, root)
	defer restorePaths()

	source := filepath.Join(AbsoluteHookyGitHooksPath, "pre-commit")
	target := filepath.Join(AbsoluteGitHooksPath, "pre-commit")
	if err := os.WriteFile(source, []byte("#!/bin/sh\necho hooky\n"), 0o750); err != nil {
		t.Fatalf("write source hook: %v", err)
	}
	if err := os.WriteFile(target, []byte("#!/bin/sh\necho legacy\n"), 0o750); err != nil {
		t.Fatalf("write target hook: %v", err)
	}

	if err := InstallHooks(InstallOptions{Force: true, Backup: false}); err != nil {
		t.Fatalf("install hooks: %v", err)
	}

	if _, err := os.Stat(target + ".hooky.bak"); !os.IsNotExist(err) {
		t.Fatalf("unexpected backup file presence, err=%v", err)
	}
}

func useTestPaths(t *testing.T, root string) func() {
	t.Helper()

	origHookyPath := AbsoluteHookyPath
	origHookyGitHooksPath := AbsoluteHookyGitHooksPath
	origGitPath := AbsoluteGitPath
	origGitHooksPath := AbsoluteGitHooksPath

	AbsoluteHookyPath = filepath.Join(root, ".hooky")
	AbsoluteHookyGitHooksPath = filepath.Join(root, ".hooky", "git-hooks")
	AbsoluteGitPath = filepath.Join(root, ".git")
	AbsoluteGitHooksPath = filepath.Join(root, ".git", "hooks")

	if err := os.MkdirAll(AbsoluteHookyGitHooksPath, 0o750); err != nil {
		t.Fatalf("create .hooky/git-hooks: %v", err)
	}
	if err := os.MkdirAll(AbsoluteGitHooksPath, 0o750); err != nil {
		t.Fatalf("create .git/hooks: %v", err)
	}

	return func() {
		AbsoluteHookyPath = origHookyPath
		AbsoluteHookyGitHooksPath = origHookyGitHooksPath
		AbsoluteGitPath = origGitPath
		AbsoluteGitHooksPath = origGitHooksPath
	}
}

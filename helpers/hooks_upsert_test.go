package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpsertGitHookOverwritesExistingHook(t *testing.T) {
	root := t.TempDir()
	restorePaths := useTestPaths(t, root)
	defer restorePaths()

	target := filepath.Join(AbsoluteHookyGitHooksPath, "pre-commit")
	if err := os.WriteFile(target, []byte("#!/bin/sh\necho old\n"), 0o750); err != nil {
		t.Fatalf("write existing hook: %v", err)
	}

	if err := UpsertGitHook("pre-commit", "go test ./..."); err != nil {
		t.Fatalf("upsert hook: %v", err)
	}

	content, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("read hook: %v", err)
	}
	if !strings.Contains(string(content), "go test ./...") {
		t.Fatalf("hook was not updated, got: %q", string(content))
	}
}

func TestCreateGitHookStillRejectsExistingHook(t *testing.T) {
	root := t.TempDir()
	restorePaths := useTestPaths(t, root)
	defer restorePaths()

	target := filepath.Join(AbsoluteHookyGitHooksPath, "pre-commit")
	if err := os.WriteFile(target, []byte("#!/bin/sh\necho old\n"), 0o750); err != nil {
		t.Fatalf("write existing hook: %v", err)
	}

	err := CreateGitHook("pre-commit", "go test ./...")
	if err == nil {
		t.Fatal("expected error for existing hook, got nil")
	}
	if !strings.Contains(err.Error(), "hook already exists") {
		t.Fatalf("unexpected error: %v", err)
	}
}

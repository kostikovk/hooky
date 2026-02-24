package helpers

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateHookyGitDirectoryUsesGitRootFromSubdirectory(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}

	restore := resetPathOverrides()
	defer restore()

	repo := t.TempDir()
	if out, err := exec.Command("git", "-C", repo, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init failed: %v: %s", err, string(out))
	}

	subdir := filepath.Join(repo, "nested", "deeper")
	if err := os.MkdirAll(subdir, 0o750); err != nil {
		t.Fatalf("mkdir subdir: %v", err)
	}

	t.Chdir(subdir)
	if err := CreateHookyGitDirectory(); err != nil {
		t.Fatalf("CreateHookyGitDirectory failed: %v", err)
	}

	expected := filepath.Join(repo, ".hooky", "hooks")
	if !dirExists(expected) {
		t.Fatalf("expected hook directory at repo root: %s", expected)
	}
}

func TestInstallHooksDoesNotCreateFakeGitDirectory(t *testing.T) {
	restore := resetPathOverrides()
	defer restore()

	workdir := t.TempDir()
	t.Chdir(workdir)

	_ = os.MkdirAll(filepath.Join(workdir, ".hooky", "hooks"), 0o750)

	err := InstallHooks(InstallOptions{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "git repository not found") {
		t.Fatalf("unexpected error: %v", err)
	}

	if dirExists(filepath.Join(workdir, ".git")) {
		t.Fatal("InstallHooks created a fake .git directory")
	}
}

func resetPathOverrides() func() {
	origHookyPath := AbsoluteHookyPath
	origHookyGitHooksPath := AbsoluteHookyGitHooksPath
	origGitPath := AbsoluteGitPath
	origGitHooksPath := AbsoluteGitHooksPath

	AbsoluteHookyPath = ""
	AbsoluteHookyGitHooksPath = ""
	AbsoluteGitPath = ""
	AbsoluteGitHooksPath = ""

	return func() {
		AbsoluteHookyPath = origHookyPath
		AbsoluteHookyGitHooksPath = origHookyGitHooksPath
		AbsoluteGitPath = origGitPath
		AbsoluteGitHooksPath = origGitHooksPath
	}
}

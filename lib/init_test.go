package lib

import (
	"errors"
	"strings"
	"testing"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func TestInitGitInitializesRepositoryWhenApproved(t *testing.T) {
	origIsGitRepository := isGitRepository
	origPromptToInitGit := promptToInitGit
	origInitGitRepository := initGitRepository
	t.Cleanup(func() {
		isGitRepository = origIsGitRepository
		promptToInitGit = origPromptToInitGit
		initGitRepository = origInitGitRepository
	})

	initCalled := false
	isGitRepository = func() bool { return false }
	promptToInitGit = func() error { return nil }
	initGitRepository = func() error {
		initCalled = true
		return nil
	}

	cmd := &cobra.Command{}
	if err := initGit(cmd); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !initCalled {
		t.Fatal("expected initGitRepository to be called")
	}
}

func TestInitGitReturnsPromptErrorWithoutInitializing(t *testing.T) {
	origIsGitRepository := isGitRepository
	origPromptToInitGit := promptToInitGit
	origInitGitRepository := initGitRepository
	t.Cleanup(func() {
		isGitRepository = origIsGitRepository
		promptToInitGit = origPromptToInitGit
		initGitRepository = origInitGitRepository
	})

	expectedErr := errors.New("declined")
	initCalled := false
	isGitRepository = func() bool { return false }
	promptToInitGit = func() error { return expectedErr }
	initGitRepository = func() error {
		initCalled = true
		return nil
	}

	cmd := &cobra.Command{}
	err := initGit(cmd)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if initCalled {
		t.Fatal("expected initGitRepository not to be called when prompt fails")
	}
}

func TestRunInitReturnsInstallError(t *testing.T) {
	origIsGitRepository := isGitRepository
	origIsHookyRepository := isHookyRepository
	origInstallHooksInGit := installHooksInGit
	t.Cleanup(func() {
		isGitRepository = origIsGitRepository
		isHookyRepository = origIsHookyRepository
		installHooksInGit = origInstallHooksInGit
	})

	isGitRepository = func() bool { return true }
	isHookyRepository = func() bool { return true }
	installErr := errors.New("install failed")
	installHooksInGit = func(options helpers.InstallOptions) error { return installErr }

	cmd := &cobra.Command{}
	cmd.Flags().Bool("force", false, "")
	cmd.Flags().Bool("backup", true, "")
	err := RunInit(cmd, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "error installing hooks") {
		t.Fatalf("expected wrapped install error, got %v", err)
	}
	if !errors.Is(err, installErr) {
		t.Fatalf("expected errors.Is(..., installErr) to be true, got %v", err)
	}
}

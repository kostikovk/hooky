package lib

import (
	"errors"
	"testing"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func TestRunAddAddsHookAndPropagatesInstallOptions(t *testing.T) {
	origUpsert := upsertGitHookInRepo
	origInstall := installHooksInGit
	t.Cleanup(func() {
		upsertGitHookInRepo = origUpsert
		installHooksInGit = origInstall
	})

	var gotHook string
	var gotCommand string
	var gotOptions helpers.InstallOptions

	upsertGitHookInRepo = func(hook, command string) error {
		gotHook = hook
		gotCommand = command
		return nil
	}
	installHooksInGit = func(options helpers.InstallOptions) error {
		gotOptions = options
		return nil
	}

	cmd := &cobra.Command{}
	cmd.Flags().Bool("force", false, "")
	cmd.Flags().Bool("backup", true, "")
	if err := cmd.Flags().Set("force", "true"); err != nil {
		t.Fatalf("set force flag: %v", err)
	}
	if err := cmd.Flags().Set("backup", "false"); err != nil {
		t.Fatalf("set backup flag: %v", err)
	}

	err := RunAdd(cmd, []string{"pre-commit", "go", "test", "./..."})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if gotHook != "pre-commit" {
		t.Fatalf("unexpected hook: got %q", gotHook)
	}
	if gotCommand != "go test ./..." {
		t.Fatalf("unexpected command: got %q", gotCommand)
	}
	if !gotOptions.Force || gotOptions.Backup {
		t.Fatalf("unexpected install options: %+v", gotOptions)
	}
}

func TestRunAddReturnsUpsertError(t *testing.T) {
	origUpsert := upsertGitHookInRepo
	t.Cleanup(func() {
		upsertGitHookInRepo = origUpsert
	})

	expectedErr := errors.New("upsert failed")
	upsertGitHookInRepo = func(hook, command string) error {
		return expectedErr
	}

	cmd := &cobra.Command{}
	cmd.Flags().Bool("force", false, "")
	cmd.Flags().Bool("backup", true, "")

	err := RunAdd(cmd, []string{"pre-commit", "go test ./..."})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

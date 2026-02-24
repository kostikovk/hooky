package lib

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunUninstallReturnsErrorAndNoSuccessMessage(t *testing.T) {
	origIsHookyRepo := isHookyRepo
	origDeleteHookyDir := deleteHookyDir
	t.Cleanup(func() {
		isHookyRepo = origIsHookyRepo
		deleteHookyDir = origDeleteHookyDir
	})

	isHookyRepo = func() bool { return true }
	deleteErr := errors.New("delete failed")
	deleteHookyDir = func() error { return deleteErr }

	cmd := &cobra.Command{}
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	err := RunUninstall(cmd, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, deleteErr) {
		t.Fatalf("expected errors.Is(..., deleteErr) true, got %v", err)
	}
	if strings.Contains(stdout.String(), "Hooky uninstalled") {
		t.Fatalf("did not expect success message, got stdout: %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr output, got %q", stderr.String())
	}
}

func TestRunUninstallPrintsSuccessOnSuccess(t *testing.T) {
	origIsHookyRepo := isHookyRepo
	origDeleteHookyDir := deleteHookyDir
	t.Cleanup(func() {
		isHookyRepo = origIsHookyRepo
		deleteHookyDir = origDeleteHookyDir
	})

	isHookyRepo = func() bool { return true }
	deleteHookyDir = func() error { return nil }

	cmd := &cobra.Command{}
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	err := RunUninstall(cmd, nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !strings.Contains(stdout.String(), "Hooky uninstalled") {
		t.Fatalf("expected success message, got stdout: %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr output, got %q", stderr.String())
	}
}

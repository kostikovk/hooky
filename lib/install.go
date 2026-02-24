package lib

import (
	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var createGitHook = helpers.CreateGitHook

func RunInstall(cmd *cobra.Command, args []string) error {
	return installHook(cmd, args[0])
}

func installHook(cmd *cobra.Command, hook string) error {
	cmd.Printf("Installing %s hook...\n", hook)

	if err := createGitHook(hook, "# go test ./..."); err != nil {
		return err
	}

	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return err
	}

	backup, err := cmd.Flags().GetBool("backup")
	if err != nil {
		return err
	}

	if err := installHooksInGit(helpers.InstallOptions{
		Force:  force,
		Backup: backup,
	}); err != nil {
		return err
	}

	cmd.Printf("Hook %s installed.\n", hook)
	return nil
}

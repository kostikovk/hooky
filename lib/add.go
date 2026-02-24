package lib

import (
	"fmt"
	"strings"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var upsertGitHookInRepo = helpers.UpsertGitHook

func RunAdd(cmd *cobra.Command, args []string) error {
	hook := args[0]
	hookCommand := strings.Join(args[1:], " ")

	cmd.Printf("Adding %s hook...\n", hook)

	if err := upsertGitHookInRepo(hook, hookCommand); err != nil {
		return err
	}

	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return fmt.Errorf("error reading --force flag: %w", err)
	}

	backup, err := cmd.Flags().GetBool("backup")
	if err != nil {
		return fmt.Errorf("error reading --backup flag: %w", err)
	}

	if err := installHooksInGit(helpers.InstallOptions{
		Force:  force,
		Backup: backup,
	}); err != nil {
		return err
	}

	cmd.Printf("Hook %s added.\n", hook)
	return nil
}

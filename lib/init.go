package lib

import (
	"fmt"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var (
	isGitRepository     = helpers.IsGitRepository
	promptToInitGit     = helpers.PromptToInitGit
	initGitRepository   = helpers.InitGit
	isHookyRepository   = helpers.IsHookyRepository
	createHookyGitDir   = helpers.CreateHookyGitDirectory
	createGitHookInRepo = helpers.CreateGitHook
	installHooksInGit   = helpers.InstallHooks
)

func RunInit(cmd *cobra.Command, args []string) error {
	cmd.Println("Initializing Hooky...")

	if err := initGit(cmd); err != nil {
		return fmt.Errorf("error initializing Git repository: %w", err)
	}

	if err := initHooky(); err != nil {
		return fmt.Errorf("error initializing Hooky repository: %w", err)
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
		return fmt.Errorf("error installing hooks: %w", err)
	}

	cmd.Println("Hooky initialized 🎉")
	return nil
}

func initGit(cmd *cobra.Command) error {
	if isGitRepository() {
		return nil
	}

	if err := promptToInitGit(); err != nil {
		return err
	}

	if err := initGitRepository(); err != nil {
		return err
	}

	cmd.Println("Git repository initialized.")
	return nil
}

func initHooky() error {
	if isHookyRepository() {
		return nil
	}

	if err := createHookyGitDir(); err != nil {
		return err
	}

	if err := createGitHookInRepo("pre-commit", "echo 'Hey 👋, Hooky!'"); err != nil {
		return err
	}

	if err := createGitHookInRepo("post-checkout", "hooky init"); err != nil {
		return err
	}

	return nil
}

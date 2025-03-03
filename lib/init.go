package lib

import (
	"os"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func RunInit(cmd *cobra.Command, args []string) {
	cmd.Println("Initializing GoHooks...")

	var err error

	err = initGit(cmd)
	if err != nil {
		cmd.Println("Error initializing Git repository.")

		os.Exit(1)
	}

	err = initHooky()
	if err != nil {
		cmd.Println("Error initializing GoHooks repository.")

		os.Exit(1)
	}

	err = helpers.InstallHooks()
	if err != nil {
		cmd.Println("Error installing hooks.")

		os.Exit(1)
	}

	cmd.Println("Hooky initialized 🎉")
}

func initGit(cmd *cobra.Command) error {
	if helpers.IsGitRepository() {
		return nil
	}

	err := helpers.PromptToInitGit()
	if err != nil {
		cmd.Println("Error initializing Git repository.")

		return err
	}

	return nil
}

func initHooky() error {
	if helpers.IsHookyRepository() {
		return nil
	}

	err := helpers.CreateHookyGitDirectory()
	if err != nil {
		return err
	}

	err = helpers.CreateGitHook("pre-commit", "echo 'Hey 👋, Hooky!'")
	if err != nil {
		return err
	}

	return nil
}

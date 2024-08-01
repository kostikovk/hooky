package lib

import (
	"os"

	"github.com/KosKosovu4/gohooks/helpers"
	"github.com/spf13/cobra"
)

func RunInit(cmd *cobra.Command, args []string) {
	cmd.Println("Initializing GoHooks...")

	var err error

	err = gitHandler(cmd)
	if err != nil {
		os.Exit(1)
	}

	err = gohooksHandler(cmd)
	if err != nil {
		os.Exit(1)
	}
}

func gitHandler(cmd *cobra.Command) error {
	if helpers.IsGitRepository() {
		cmd.Println("Git already initialized.")

		return nil
	}

	err := helpers.PromptToInitGit()
	if err != nil {
		cmd.Println("Error initializing Git repository.")

		return err
	}

	cmd.Println("Git repository initialized.")

	return nil
}

func gohooksHandler(cmd *cobra.Command) error {
	// Check if GoHooks repository already exists
	if helpers.IsGoHooksRepository() {
		cmd.Println("GoHooks already initialized.")

		return nil
	}

	// Create GoHooks repository
	err := helpers.CreateGoHooksGitDirectory()
	if err != nil {
		cmd.Println("Error creating GoHooks repository.")

		return err
	}

	cmd.Println("GoHooks repository initialized.")

	return nil
}

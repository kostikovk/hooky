package lib

import (
	"os"

	"github.com/2kse/gohooks/helpers"
	"github.com/spf13/cobra"
)

func RunInit(cmd *cobra.Command, args []string) {
	cmd.Println("Initializing GoHooks...")

	var err error

	// Initialize or ask to initialize Git repository
	err = initGit(cmd)
	if err != nil {
		cmd.Println("Error initializing Git repository.")

		os.Exit(1)
	} else {
		cmd.Println("Git repository initialized.")
	}

	// Initialize GoHooks repository
	err = initGoHooks()
	if err != nil {
		cmd.Println("Error initializing GoHooks repository.")

		os.Exit(1)
	} else {
		cmd.Println("GoHooks repository initialized.")
	}

	// Copy Git hooks to GoHooks repository after prompting the user to do so
	// This is optional and can be skipped by the user if they don't want to copy Git hooks
	copyGitHooksToGoHooks()
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

func initGoHooks() error {
	// Check if GoHooks repository already exists
	if helpers.IsGoHooksRepository() {
		return nil
	}

	// Create GoHooks repository
	err := helpers.CreateGoHooksGitDirectory()
	if err != nil {
		return err
	}

	return nil
}

func copyGitHooksToGoHooks() error {
	if !helpers.HasGitHooks() {
		return nil
	}

	return helpers.PromptToCopyGitHooksToGoHooks()
}

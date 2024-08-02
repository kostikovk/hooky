package lib

import (
	"os"

	"github.com/2kse/gohooks/helpers"
	"github.com/spf13/cobra"
)

// RunInit initializes GoHooks.
func RunInit(cmd *cobra.Command, args []string) {
	cmd.Println("Initializing GoHooks...")

	var err error

	// Initialize or ask to initialize Git repository
	err = initGit(cmd)
	if err != nil {
		cmd.Println("Error initializing Git repository.")

		os.Exit(1)
	}

	// Initialize GoHooks repository
	err = initGoHooks()
	if err != nil {
		cmd.Println("Error initializing GoHooks repository.")

		os.Exit(1)
	}

	err = helpers.InstallHooks()
	if err != nil {
		cmd.Println("Error installing hooks.")

		os.Exit(1)
	}

	cmd.Println("GoHooks initialized.")
}

// Initialize or ask to initialize Git repository.
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

// Initialize GoHooks repository.
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

	// Create pre-commit hook
	err = helpers.CreateGitHook("pre-commit", "# go test ./...")
	if err != nil {
		return err
	}

	return nil
}

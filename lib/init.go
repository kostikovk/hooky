package lib

import (
	"os"
	"sync"

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

	var wg sync.WaitGroup
	errors := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := helpers.CreateGitHook("pre-commit", "echo 'Hey 👋, Hooky!'"); err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := helpers.CreateGitHook("post-checkout", "hooky init"); err != nil {
			errors <- err
		}
	}()

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			return err
		}
	}

	return nil
}

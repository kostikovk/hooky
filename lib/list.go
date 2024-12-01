package lib

import (
	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

// Whant to implement:
// 1. Show all available hooks
// 2. Show all hooks that are installed
// 3. Show all hooks that are not installed
// 4. Ask a user about installing a hook

func RunList(cmd *cobra.Command, args []string) {
	installed, _ := cmd.Flags().GetBool("installed")

	if installed {
		showListOfInstalledHooks(cmd)
	} else {
		showListOfAvailableHooks(cmd)
	}
}

func showListOfAvailableHooks(cmd *cobra.Command) {
	cmd.Println("Git Hooks:")

	for i, hook := range helpers.GitHooks {
		cmd.Printf("%d. %s\n", i, hook)
	}
}

func showListOfInstalledHooks(cmd *cobra.Command) {
	cmd.Println("Installed Git Hooks:")

	var installedHooks []string
	var err error
	installedHooks, err = helpers.ListOfInstalledGitHooks()
	if err != nil {
		cmd.PrintErr("Error listing installed hooks.")
	}

	for i, hook := range installedHooks {
		cmd.Printf("%d. %s\n", i, hook)
	}
}
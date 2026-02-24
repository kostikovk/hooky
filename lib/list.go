package lib

import (
	"fmt"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var listInstalledHooks = helpers.ListOfInstalledGitHooks

func RunList(cmd *cobra.Command, args []string) error {
	installed, _ := cmd.Flags().GetBool("installed")

	if installed {
		return showListOfInstalledHooks(cmd)
	}
	showListOfAvailableHooks(cmd)
	return nil
}

func showListOfAvailableHooks(cmd *cobra.Command) {
	cmd.Println("Git Hooks:")

	for i, hook := range helpers.GitHooks {
		cmd.Printf("%d. %s\n", i, hook)
	}
}

func showListOfInstalledHooks(cmd *cobra.Command) error {
	cmd.Println("Installed Git Hooks:")

	installedHooks, err := listInstalledHooks()
	if err != nil {
		return fmt.Errorf("failed to list installed hooks: %w", err)
	}

	for i, hook := range installedHooks {
		cmd.Printf("%d. %s\n", i, hook)
	}
	return nil
}

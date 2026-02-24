package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [hook] [command]",
	Short: "Add or update a Git hook command",
	Long:  `Add or update a Git hook script in .hooky/git-hooks, then sync it into .git/hooks.`,
	Args:  cobra.MinimumNArgs(2),
	RunE:  lib.RunAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Bool("force", false, "Replace existing hooks in .git/hooks when conflicts are found")
	addCmd.Flags().Bool("backup", true, "Backup existing hooks before replacing (used with --force)")
}

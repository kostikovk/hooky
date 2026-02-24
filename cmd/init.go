package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Hooky CLI",
	Long:  `Init Hooky CLI...`,
	RunE:  lib.RunInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().Bool("force", false, "Replace existing hooks in .git/hooks when conflicts are found")
	initCmd.Flags().Bool("backup", true, "Backup existing hooks before replacing (used with --force)")
}

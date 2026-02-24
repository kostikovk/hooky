package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available hooks",
	Long:  `List all available hooks.`,
	RunE:  lib.RunList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("installed", false, "Show only installed hooks")
}

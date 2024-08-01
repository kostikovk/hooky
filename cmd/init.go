package cmd

import (
	"github.com/KosKosovu4/gohooks/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init GoHooks CLI",
	Long:  `Init GoHooks CLI...`,
	Run:   lib.RunInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

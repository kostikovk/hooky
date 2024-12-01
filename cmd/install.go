package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [hook]",
	Short: "Install a specific Git hook",
	Long:  `Install a specific Git hook, such as pre-commit or pre-push, and set up the necessary scripts to execute the hook logic.`,
	Args:  cobra.ExactArgs(1),
	Run:   lib.RunInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

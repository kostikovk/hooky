package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall Hooky CLI",
	Long:  `Uninstall Hooky CLI...`,
	Run:   lib.RunUninstall,
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

package lib

import (
	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func RunUninstall(cmd *cobra.Command, args []string) {
	err := hookyUninstallHandler(cmd)
	if err != nil {
		cmd.Println("Error uninstalling GoHooks.")
	}

	cmd.Println("Hooky uninstalled 🥺")
}

func hookyUninstallHandler(cmd *cobra.Command) error {
	if !helpers.IsHookyRepository() {
		return nil
	}

	err := helpers.DeleteHookyDirectory()
	if err != nil {
		cmd.Println("Error deleting GoHooks repository.")

		return err
	}

	return nil
}

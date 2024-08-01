package lib

import (
	"github.com/2kse/gohooks/helpers"
	"github.com/spf13/cobra"
)

func RunUninstall(cmd *cobra.Command, args []string) {
	err := gohooksUninstallHandler(cmd)
	if err != nil {
		cmd.Println("Error uninstalling GoHooks.")
	}

	cmd.Println("GoHooks uninstalled.")
}

func gohooksUninstallHandler(cmd *cobra.Command) error {
	if !helpers.IsGoHooksRepository() {
		return nil
	}

	err := helpers.DeleteGoHooksDirectory()
	if err != nil {
		cmd.Println("Error deleting GoHooks repository.")

		return err
	}

	return nil
}

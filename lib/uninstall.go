package lib

import (
	"github.com/KosKosovu4/gohooks/helpers"
	"github.com/spf13/cobra"
)

func RunUninstall(cmd *cobra.Command, args []string) {
	cmd.Println("Uninstalling GoHooks...")

	err := gohooksUninstallHandler(cmd)
	if err != nil {
		cmd.Println("Error uninstalling GoHooks.")
	}

	cmd.Println("GoHooks uninstalled.")
}

func gohooksUninstallHandler(cmd *cobra.Command) error {
	if !helpers.IsGoHooksRepository() {
		cmd.Println("GoHooks repository does not exist.")

		return nil
	}

	err := helpers.DeleteGoHooksDirectory()
	if err != nil {
		cmd.Println("Error deleting GoHooks repository.")

		return err
	}

	return nil
}

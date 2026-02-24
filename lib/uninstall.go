package lib

import (
	"fmt"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

var (
	isHookyRepo    = helpers.IsHookyRepository
	deleteHookyDir = helpers.DeleteHookyDirectory
)

func RunUninstall(cmd *cobra.Command, args []string) error {
	if err := hookyUninstallHandler(); err != nil {
		return fmt.Errorf("uninstall failed: %w", err)
	}

	cmd.Println("Hooky uninstalled 🥺")
	return nil
}

func hookyUninstallHandler() error {
	if !isHookyRepo() {
		return nil
	}

	if err := deleteHookyDir(); err != nil {
		return fmt.Errorf("error deleting Hooky repository: %w", err)
	}

	return nil
}

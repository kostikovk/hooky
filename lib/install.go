package lib

import (
	"fmt"
	"os"

	"github.com/kostikovk/hooky/helpers"
	"github.com/spf13/cobra"
)

func RunInstall(cmd *cobra.Command, args []string) {
	installHook(cmd, args[0])
}

func installHook(cmd *cobra.Command, hook string) {
	fmt.Printf("Installing %s hook...\n", hook)

	err := helpers.CreateGitHook(hook, "# go test ./...")
	if err != nil {
		cmd.PrintErr(err)

		os.Exit(1)
	}

	cmd.Printf("Hook %s installed.", hook)
}

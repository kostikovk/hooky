package cmd

import (
	"github.com/spf13/cobra"
)

// Version is the version of the CLI.
// It can be overridden at build time:
// go build -ldflags="-X github.com/kostikovk/hooky/cmd.Version=vX.Y.Z"
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println(Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

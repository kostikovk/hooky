package cmd

import (
	"github.com/spf13/cobra"
)

// Version is the version of the CLI.
const Version string = "v1.0.1"

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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is the version of the CLI.
var Version string = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Version information")

	cobra.OnInitialize(checkVersionFlag)
}

func checkVersionFlag() {
	if v, err := rootCmd.Flags().GetBool("version"); err == nil && v {
		fmt.Println(Version)
		os.Exit(0)
	}
}

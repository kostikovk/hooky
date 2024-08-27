package cmd

import (
	"fmt"

	"github.com/kostikovk/gohooks/helpers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available hooks",
	Long:  `List all available hooks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List all available hooks:")

		// todo: need to implement this function
		for i, hook := range helpers.GitHooks {
			fmt.Printf("%d. %s\n", i+1, hook)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

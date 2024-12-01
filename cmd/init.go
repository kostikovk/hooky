package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Hooky CLI",
	Long:  `Init Hooky CLI...`,
	Run:   lib.RunInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

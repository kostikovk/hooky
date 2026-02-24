package cmd

import (
	"github.com/kostikovk/hooky/lib"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Validate Hooky installation and hook wiring",
	Long:  `Run diagnostics for Hooky state, including repository layout and .git/hooks symlink wiring.`,
	RunE:  lib.RunDoctor,
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

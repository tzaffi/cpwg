/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tzaffi/cpwg/chp2/pkg/grep"
)

// grepfilesCmd represents the grepfiles command
var grepfilesCmd = &cobra.Command{
	Use:   "grepfiles",
	Short: "Search for a pattern in a list of files and print the matches",
	Run: func(cmd *cobra.Command, args []string) {
		if err := grep.Grep(args[1:], args[0]); err != nil {
			cmd.PrintErr(err, "\n")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(grepfilesCmd)
}

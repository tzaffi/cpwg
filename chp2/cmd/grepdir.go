/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tzaffi/cpwg/chp2/pkg/grepdir"
)

// grepdirCmd represents the grepdir command
var grepdirCmd = &cobra.Command{
	Use:   "grepdir",
	Short: "Search for a patterin in a directory and print the matches",
	Run: func(cmd *cobra.Command, args []string) {
		if err := grepdir.GrepDir(args[1], args[0]); err != nil {
			cmd.PrintErr(err, "\n")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(grepdirCmd)
}

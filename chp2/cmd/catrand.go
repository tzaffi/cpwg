/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tzaffi/cpwg/chp2/pkg/cat"
)

// catrandCmd represents the catrand command
var catrandCmd = &cobra.Command{
	Use:   "catrand",
	Short: "Concatenates provided files, but not necessarily in the order provided.",
	Long:  `This is also Ex. 1 of Chapter 2.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cat.CatRand(args); err != nil {
			cmd.PrintErr(err, "\n")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(catrandCmd)
}

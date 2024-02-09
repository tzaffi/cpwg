/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tzaffi/cpwg/chp2/pkg/greprec"
)

// greprecCmd represents the greprec command
var greprecCmd = &cobra.Command{
	Use:   "greprec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := greprec.GrepRec(args[1], args[0]); err != nil {
			cmd.PrintErr(err, "\n")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(greprecCmd)
}

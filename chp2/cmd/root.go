/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tzaffi/cpwg/chp2/pkg/cat"
	"github.com/tzaffi/cpwg/chp2/pkg/grep"
	"github.com/tzaffi/cpwg/chp2/pkg/grepdir"
)

var files []string
var dir string
var pattern string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chp2",
	Short: "General purpose text file processing switched via flags or subcommands",
	// Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(files) == 0 && len(dir) == 0 {
			cmd.PrintErr("No files or dir provided\n")
			cmd.Help()
			os.Exit(1)
		}
		if len(files) > 0 && len(dir) > 0 {
			cmd.PrintErr("Cannot use both files and dir flags\n")
			cmd.Help()
			os.Exit(1)
		}
		// if pattern was provided, use it to call grep.Grep
		// otherwise, call cat.CatRand
		if pattern != "" {
			if len(files) > 0 {
				if err := grep.Grep(files, pattern); err != nil {
					cmd.PrintErr(err, "\n")
					os.Exit(1)
				}
			} else if err := grepdir.GrepDir(dir, pattern); err != nil {
				cmd.PrintErr(err, "\n")
				os.Exit(1)
			}
			return
		}

		if err := cat.CatRand(files); err != nil {
			cmd.PrintErr(err, "\n")
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&files, "files", "f", []string{}, "Files for processing")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "Directory for processing")
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Pattern to search for in the files")
}

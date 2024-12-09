/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var shift int
var outputPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "caesar-cipher",
	Short: "Encrypt and decrypt messages using a Caesar cipher",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if shift < 1 || shift > 25 {
			fmt.Printf("Invalid shift key: %d\n", shift)
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
	rootCmd.PersistentFlags().IntVarP(&shift, "shift", "s", 8, "Key used to encrypt the message (between 1 and 25)")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "Path to the output file (Prints to stdout if not specified)")
}

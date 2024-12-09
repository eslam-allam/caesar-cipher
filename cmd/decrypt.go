/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/eslam-allam/caesar-cipher/internal"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt a message using a shift key",
    Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var data string
		var err error
		if len(args) == 0 {
			data, err = internal.GetStdin()
		} else {
			data, err = internal.ReadFile(args[0])
		}
		if err != nil {
			if errors.Is(err, internal.ErrStdEmpty) {
				fmt.Print("No file specified\n")
				os.Exit(1)
			}
			fmt.Printf("Failed to read data: %v\n", err)
			os.Exit(1)
		}
		decrypted := internal.Decrypt(data, shift)

		if outputPath != "" {
			err = internal.WriteFile(outputPath, decrypted, "decrypted.txt")
			if err != nil {
				fmt.Printf("Failed to write output file: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}
		fmt.Println(decrypted)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

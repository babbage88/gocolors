/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/babbage88/gocolors/internal/header"
	"github.com/spf13/cobra"
)

// testCmdCmd represents the testCmd command
var testCmdCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test := header.PrintHeaderMessage("This is test yes.", "#")
		fmt.Printf("%s\n", test)
	},
}

func init() {
	rootCmd.AddCommand(testCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

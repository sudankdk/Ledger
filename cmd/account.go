/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Long: `can be used to create accounts, list accounts, and get account details.
Example usage:
  ledger account create "Checking Account"
  ledger account list
  ledger account get 1`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("account called")
	// },
}

func init() {
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

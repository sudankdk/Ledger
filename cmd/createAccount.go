/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sudankdk/ledger/internal/db/service"
)

var Service *service.SQLService

// createAccountCmd represents the createAccount command
var createAccountCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Long:  `Creates a new account with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: ledger account create <account_name>")
			return
		}
		accountName := args[0]
		fmt.Printf("Creating account: %s\n", accountName)

		account, err := Service.CreateAccount(context.Background(), accountName)
		if err != nil {
			fmt.Printf("Error creating account: %v\n", err)
			return
		}
		fmt.Printf("Account created successfully: %v\n", account)
	},
}

var listAccountsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	Long:  `Lists all accounts in the ledger.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing accounts...")

		accounts, err := Service.ListAccounts(context.Background())
		if err != nil {
			fmt.Printf("Error listing accounts: %v\n", err)
			return
		}
		for _, account := range accounts {
			fmt.Printf("- %s (ID: %d)\n", account.Name, account.ID)
		}
	},
}

func init() {
	accountCmd.AddCommand(createAccountCmd)
	accountCmd.AddCommand(listAccountsCmd)

}

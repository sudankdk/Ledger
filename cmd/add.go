/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a transaction",
	Long:  `Add a new transaction to the ledger.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var amount float64
var description string

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "The amount of the transaction")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "The description of the transaction")

	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("description")
}

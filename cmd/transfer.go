package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sudankdk/ledger/internal/db/service"
)

var transferService *service.SQLService

// transferCmd performs a double-entry transfer between two accounts by name.
var transferCmd = &cobra.Command{
	Use:   "transfer <amount> <from_account_name> <to_account_name> [description]",
	Short: "Transfer amount between accounts",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		amtStr := args[0]
		from := args[1]
		to := args[2]
		desc := ""
		if len(args) > 3 {
			desc = args[3]
		}

		amt, err := strconv.ParseFloat(amtStr, 64)
		if err != nil {
			fmt.Printf("invalid amount: %v\n", err)
			return
		}

		// resolve accounts by name
		fromAcc, err := Service.GetAccountByName(context.Background(), from)
		if err != nil {
			fmt.Printf("failed to find from account: %v\n", err)
			return
		}
		toAcc, err := Service.GetAccountByName(context.Background(), to)
		if err != nil {
			fmt.Printf("failed to find to account: %v\n", err)
			return
		}

		tx, err := Service.DoTransaction(context.Background(), amt, desc, fromAcc.ID, toAcc.ID)
		if err != nil {
			fmt.Printf("transaction failed: %v\n", err)
			return
		}
		fmt.Printf("transaction created: %+v\n", tx)
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
}

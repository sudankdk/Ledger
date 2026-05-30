package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// AccountService defines the minimal interface the CLI expects.
type AccountService interface {
    Create(name string) error
}

// accountService should be injected by application wiring (set from main).
var accountService AccountService

var accountCreateCmd = &cobra.Command{
    Use:   "create",
    Short: "Create account",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]
        if accountService == nil {
            return fmt.Errorf("accountService is not initialized; set cmd.accountService in main")
        }
        return accountService.Create(name)
    },
}

func init() {
    accountCmd.AddCommand(accountCreateCmd)
}

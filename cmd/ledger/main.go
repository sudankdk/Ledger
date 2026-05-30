package main

import (
	"fmt"

	"github.com/sudankdk/ledger/cmd"
)

func main() {
	fmt.Println("Welcome to the Ledger CLI!")
	cmd.Execute()
}

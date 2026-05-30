package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sudankdk/ledger/cmd"
	dbservice "github.com/sudankdk/ledger/internal/db/service"
	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Welcome to the Ledger CLI!")

	db, err := sql.Open("sqlite", "ledger.db")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping failed: %v", err)
	}

	svc := dbservice.NewSQLService(db)
	cmd.Service = svc

	if len(os.Args) > 1 {
		if err := cmd.Execute(); err != nil {
			log.Fatalf("command error: %v", err)
		}
		return
	}

	// Interactive REPL
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("ledger> ")
		if !scanner.Scan() {
			fmt.Println()
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			break
		}
		args := strings.Fields(line)
		if err := cmd.ExecuteArgs(args); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/peterh/liner"
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

	// Interactive REPL with line editing & history
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	histFile := filepath.Join(".", ".ledger_history")
	if f, err := os.Open(histFile); err == nil {
		_, _ = line.ReadHistory(f)
		f.Close()
	}

	completer := func(input string) (c []string) {
		cmds := []string{"account create", "account list", "transfer", "help", "exit", "quit"}
		for _, s := range cmds {
			if len(input) == 0 || (len(input) <= len(s) && s[:len(input)] == input) || (len(input) > 0 && s[:len(input)] == input) {
				c = append(c, s)
			}
		}
		return
	}
	line.SetCompleter(func(lineStr string) (c []string) { return completer(lineStr) })

	for {
		input, err := line.Prompt("ledger> ")
		if err != nil {
			break
		}
		input = string(input)
		if input == "" {
			continue
		}
		if input == "exit" || input == "quit" || input == "q" {
			break
		}
		line.AppendHistory(input)
		args := splitFields(input)
		if err := cmd.ExecuteArgs(args); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}

	if f, err := os.Create(histFile); err == nil {
		_, _ = line.WriteHistory(f)
		f.Close()
	}
}

// splitFields is like strings.Fields but keeps quoted substrings together.
func splitFields(s string) []string {
	var out []string
	cur := ""
	inQuotes := false
	quoteChar := '\''
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if inQuotes {
			if ch == byte(quoteChar) {
				inQuotes = false
				out = append(out, cur)
				cur = ""
			} else {
				cur += string(ch)
			}
			continue
		}
		if ch == ' ' || ch == '\t' {
			if cur != "" {
				out = append(out, cur)
				cur = ""
			}
			continue
		}
		if ch == '\'' || ch == '"' {
			inQuotes = true
			quoteChar = rune(ch)
			continue
		}
		cur += string(ch)
	}
	if cur != "" {
		out = append(out, cur)
	}
	return out
}

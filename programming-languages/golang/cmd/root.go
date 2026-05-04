package cmd

import (
	"bufio"
	"fmt"
	"learning-golang/internal/store"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Store store.Store

var rootCmd = &cobra.Command{
	Use:   "stok",
	Short: "stok is a kv store",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("stok ready. type 'exit' to quit.")
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			if line == "exit" || line == "quit" {
				return
			}
			parts := splitArgs(line)
			subCmd, subArgs, err := cmd.Find(parts)
			if err != nil || subCmd == cmd {
				fmt.Fprintln(os.Stderr, "unknown command:", parts[0])
				continue
			}
			if err := subCmd.ValidateArgs(subArgs); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			subCmd.Run(subCmd, subArgs)
		}
	},
}

func splitArgs(line string) []string {
	var args []string
	var current strings.Builder
	inQuote := false
	var quoteChar rune

	for _, ch := range line {
		switch {
		case inQuote && ch == quoteChar:
			inQuote = false
		case !inQuote && (ch == '"' || ch == '\''):
			inQuote = true
			quoteChar = ch
		case !inQuote && ch == ' ':
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args
}

func Execute(s store.Store) {
	Store = s
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

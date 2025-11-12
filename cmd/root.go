package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/nghSia/Mini-CRM/internal/storage"
	"github.com/spf13/cobra"
)

var (
	store  storage.Storer
	reader *bufio.Reader

	rootCmd = &cobra.Command{
		Use:   "gomincrm",
		Short: "A CLI tool to manage contacts in Mini-CRM",
		Long: `gomincrm is a command-line interface application for managing contacts.
It provides commands to add, list, view, update, and delete contacts.
You can use it in interactive mode or with flags for quick operations.

Examples:
  gomincrm add              # Interactive mode to add a contact
  gomincrm add -n "John" -e "john@example.com"  # Add with flags
  gomincrm list             # List all contacts
  gomincrm get 1            # Get contact with ID 1
  gomincrm update 1 -n "Jane"  # Update contact name
  gomincrm delete 1         # Delete contact with ID 1`,
		Run: func(cmd *cobra.Command, args []string) {
			app.Run(store)
		},
	}
)

func Execute() {
	store = storage.NewGORMStore()
	reader = bufio.NewReader(os.Stdin)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func GetStore() storage.Storer { return store }
func GetReader() *bufio.Reader { return reader }

func intit() {}

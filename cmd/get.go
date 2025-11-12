package cmd

import (
	"fmt"
	"strconv"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get a contact by ID",
	Long: `Retrieve and display detailed information about a specific contact.

Provide the contact ID as an argument, or run in interactive mode
to be prompted for the ID.

Examples:
  gomincrm get 1     # Get contact with ID 1
  gomincrm get       # Interactive mode - will prompt for ID`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := GetStore()
		reader := GetReader()

		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err != nil || id <= 0 {
				fmt.Println("❌ Please enter a vid ID.")
				return nil
			}

			contact, err := store.GetById(id)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}

			fmt.Println("📋 Contact's informations :")
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.Id, contact.Name, contact.Email)
			return nil
		}

		app.HandleGetContactByID(reader, store)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

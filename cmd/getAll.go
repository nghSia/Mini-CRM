package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts",
	Long: `Display a list of all contacts in the Mini-CRM system.

This command shows all contacts with their ID, name, and email address.
If no contacts exist, you'll receive a message indicating the list is empty.

Examples:
  gomincrm list     # Display all contacts`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := GetStore()

		contacts, err := store.GetAll()
		if err != nil {
			return err
		}

		fmt.Println("📋 Users list:")
		for _, contact := range contacts {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.Id, contact.Name, contact.Email)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}

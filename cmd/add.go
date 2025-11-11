package cmd

import (
	"fmt"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/nghSia/Mini-CRM/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact to the CRM",
	Long: `Add a new contact to the Mini-CRM system.

You can use this command in two ways:
1. Interactive mode: Simply run 'gomincrm add' and follow the prompts
2. With flags: Provide name and email directly using -n and -e flags

Examples:
  gomincrm add                           # Interactive mode
  gomincrm add -n "John Doe" -e "john@example.com"  # Using flags`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := GetStore()
		reader := GetReader()

		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name != "" && email != "" {
			contact := &storage.Contact{Name: name, Email: email}
			if err := store.Add(contact); err != nil {
				return err
			}
			fmt.Printf("✅ Contact added: %s (%s)\n", name, email)
			return nil
		}

		app.HandleAddContact(reader, store)
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Contact Name")
	addCmd.Flags().StringP("email", "e", "", "Contact Email")

	rootCmd.AddCommand(addCmd)
}

package cmd

import (
	"fmt"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an existing contact",
	Long: `Update an existing contact's information in the Mini-CRM system.

You can update either the name, email, or both fields of a contact.
Use flags to update specific fields, or run without flags for interactive mode.

Examples:
  gomincrm update                           # Interactive mode
  gomincrm update -i 1 -n "Jane Doe"        # Update name only
  gomincrm update -i 1 -e "jane@example.com"  # Update email only
  gomincrm update -i 1 -n "Jane" -e "jane@example.com"  # Update both`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := GetStore()
		reader := GetReader()

		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if id > 0 && (name != "" || email != "") {
			if err := store.Update(id, name, email); err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println("✅ Contact updated successfully")
			return nil
		}

		app.HandleUpdateContact(reader, store)
		return nil
	},
}

func init() {
	updateCmd.Flags().IntP("id", "i", 0, "ID of the contact to update")
	updateCmd.Flags().StringP("name", "n", "", "New name of the contact")
	updateCmd.Flags().StringP("email", "e", "", "New email of the contact")

	rootCmd.AddCommand(updateCmd)
}

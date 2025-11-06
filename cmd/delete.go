package cmd

import (
	"fmt"
	"strconv"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a contact by ID",
	Long: `Delete a contact from the Mini-CRM system using its ID.
You can provide the ID as an argument or run in interactive mode
to be prompted for the ID.

Examples:
  gomincrm delete 1     # Delete contact with ID 1
  gomincrm delete       # Interactive mode - will prompt for ID`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := GetStore()
		reader := GetReader()

		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err != nil || id <= 0 {
				fmt.Println("❌ Please enter a vid ID.")
				return nil
			}

			if err := store.Delete(id); err != nil {
				return err
			}
			fmt.Println("✅ Contact deleted successfully")
			return nil
		}

		app.HandleDeleteContact(reader, store)
		return nil
	},
}

func init() {
	deleteCmd.Flags().IntP("id", "i", 0, "Contact ID to delete")

	rootCmd.AddCommand(deleteCmd)
}

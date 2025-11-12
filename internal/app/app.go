package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nghSia/Mini-CRM/internal/storage"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n========== Mini-CRM Menu ==========")
		fmt.Println("1) : Add a contact")
		fmt.Println("2) : List all contacts")
		fmt.Println("3) : Show contact information")
		fmt.Println("4) : Update a contact")
		fmt.Println("5) : Delete a contact")
		fmt.Print("6) Quit application")
		fmt.Println("\n===================================")

		choice := readInt(reader, "Enter your choice: ")

		switch choice {
		case 1:
			HandleAddContact(reader, store)
		case 2:
			HandleGetAllContact(store)
		case 3:
			HandleGetContactByID(reader, store)
		case 4:
			HandleUpdateContact(reader, store)
		case 5:
			HandleDeleteContact(reader, store)
		case 6:
			fmt.Println("Closing...")
			return
		default:
			fmt.Println(choice)
		}
	}
}

func HandleAddContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter username: ")
	i_username, _ := reader.ReadString('\n')
	i_username = strings.TrimSpace(i_username)

	fmt.Print("Enter user email: ")
	i_userMail, _ := reader.ReadString('\n')
	i_userMail = strings.TrimSpace(i_userMail)

	if i_username == "" || i_userMail == "" {
		fmt.Println("❌ Invalid values, please enter a valid name and email.")
		return
	}

	newUser := &storage.Contact{Name: i_username, Email: i_userMail}

	err := store.Add(newUser)
	if err != nil {
		fmt.Printf("❌ Error adding contact: %v\n", err)
		return
	}
	fmt.Printf("✅ Contact added successfully: %s (%s)\n", i_username, i_userMail)
}

func HandleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter user ID to delete: ")
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	if err != nil {
		fmt.Printf("❌ Invalid ID, please enter a valid ID.")
		return
	}

	err = store.Delete(i_index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("✅ Contact deleted successfully")
}

func HandleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter user ID to update: ")
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	contactsLength, err := store.GetAll()

	if len(contactsLength) < i_index {
		fmt.Println("❌ No matching contact found with the given ID.")
		return
	}

	fmt.Print("Enter new name: ")
	i_name, _ := reader.ReadString('\n')
	i_name = strings.TrimSpace(i_name)

	fmt.Print("Enter new email: ")
	i_email, _ := reader.ReadString('\n')
	i_email = strings.TrimSpace(i_email)

	err = store.Update(i_index, i_name, i_email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("✅ Contact updated successfully")
}

func HandleGetAllContact(store storage.Storer) {
	contacts, err := store.GetAll()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, contact := range contacts {
		fmt.Println("📋 Users list:")
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", contact.Id, contact.Name, contact.Email)
	}
}

func HandleGetContactByID(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter user ID to display: ")
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	if err != nil {
		fmt.Println("❌ Invalid ID, please enter a valid integer.")
		return
	}

	contact, err := store.GetById(i_index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ID: %d | Name: %s | Email: %s\n", contact.Id, contact.Name, contact.Email)
}

func readInt(reader *bufio.Reader, inputMessage string) int {
	for {
		fmt.Print(inputMessage)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			fmt.Println("❌ Please enter a valid value")
			continue
		}

		value, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("❌ Please enter a valid integer.")
			continue
		}

		if value < 1 || value > 6 {
			fmt.Println("❌ Please enter a number between 1 and 6.")
			continue
		}
		return value
	}
}

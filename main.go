package main

import (
	"flag"
	"fmt"

	"github.com/nghSia/Mini-CRM/contact"
	"github.com/nghSia/Mini-CRM/crudcontact"
	"github.com/nghSia/Mini-CRM/menu"
)

func main() {
	name := flag.String("name", "", "Nom du contact à ajouter")
	email := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()

	if *name != "" || *email != "" {
		if *name == "" || *email == "" {
			fmt.Println("❌ Vous devez fournir un nom ET un email avec -name et -email")
		} else {
			newUser := contact.Contact{Name: *name, Email: *email}
			crudcontact.AddContactToList(newUser)
			fmt.Printf("✅ Contact ajouté : %s (%s)\n\n", *name, *email)
		}
	}

	menu.DisplayMenu()
}

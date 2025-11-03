package main

import (
	"flag"
	"fmt"

	crudcontact "github.com/nghSia/Mini-CRM/crudContact"
	"github.com/nghSia/Mini-CRM/menu"
	"github.com/nghSia/Mini-CRM/user"
)

func main() {
	name := flag.String("name", "", "Nom du contact à ajouter")
	email := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()

	if *name != "" || *email != "" {
		if *name == "" || *email == "" {
			fmt.Println("❌ Vous devez fournir un nom ET un email avec -name et -email")
		} else {
			newUser := user.User{Name: *name, Email: *email}
			crudcontact.AddUser(newUser)
			fmt.Printf("✅ Contact ajouté : %s (%s)\n\n", *name, *email)
		}
	}

	menu.DisplayMenu()
}

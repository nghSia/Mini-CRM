package main

import (
	"flag"
	"fmt"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/nghSia/Mini-CRM/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()

	name := flag.String("name", "", "Nom du contact à ajouter")
	email := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()

	if *name != "" || *email != "" {
		if *name == "" || *email == "" {
			fmt.Println("❌ Vous devez fournir un nom ET un email avec -name et -email")
		} else {
			newUser := &storage.Contact{Name: *name, Email: *email}
			store.Add(newUser)
			fmt.Printf("✅ Contact ajouté : %s (%s)\n\n", *name, *email)
		}
	}

	app.Run(store)
}

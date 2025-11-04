package crudcontact

import (
	"fmt"

	"github.com/nghSia/Mini-CRM/contact"
)

// une seule instance en mémoire, tout le monde modifie la même.
var ListUsers = make(map[int]*contact.Contact)

var nextID int

func GetUsers() {
	if len(ListUsers) == 0 {
		fmt.Println("Aucun contact pour l’instant.")
		return
	}

	fmt.Println("\n📋 Liste des utilisateurs :")
	for _, user := range ListUsers {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", user.Id, user.Name, user.Email)
	}
}

func AddContactToList(p_user contact.Contact) {
	nextID++
	p_user.Id = nextID
	userToAdd, err := p_user.Add()
	if err != nil {
		fmt.Println("❌", err)
		return
	}
	ListUsers[nextID] = userToAdd
}

func UpdateContactList(id int, p_user contact.Contact) {
	foundContact, exists := ListUsers[id]

	if !exists {
		fmt.Printf("❌ Aucun utilisateur trouvé avec l’ID %d\n", id)
		return
	}

	_, err := foundContact.Update(p_user.Name, p_user.Email)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Printf("✅ Utilisateur avec l’ID %d mis à jour avec succès\n", id)
}

func DeleteUser(p_id int) {
	_, exists := ListUsers[p_id]

	if !exists {
		fmt.Printf("❌ Aucun utilisateur trouvé avec l’ID %d\n", p_id)
		return
	}

	delete(ListUsers, p_id)
	fmt.Printf("✅ Utilisateur avec l’ID %d supprimé avec succès\n", p_id)
}

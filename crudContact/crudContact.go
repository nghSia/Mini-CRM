package crudcontact

import (
	"fmt"

	"github.com/nghSia/Mini-CRM/user"
)

var ListUsers = make(map[int]user.User)

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

func AddUser(p_user user.User) {
	nextID++
	p_user.Id = nextID
	ListUsers[p_user.Id] = p_user
}

func UpdateUser(p_id int, p_user user.User) {
	_, exists := ListUsers[p_id]

	if !exists {
		fmt.Printf("❌ Aucun utilisateur trouvé avec l’ID %d\n", p_id)
		return
	}

	p_user.Id = p_id
	ListUsers[p_id] = p_user

	fmt.Printf("✅ Utilisateur avec l’ID %d mis à jour avec succès\n", p_id)
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

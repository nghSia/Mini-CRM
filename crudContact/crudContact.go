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

	fmt.Println("\n📋 Liste des contacts :")
	for _, user := range ListUsers {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", user.Id, user.Name, user.Email)
	}
}

func Add(p_user user.User) {
	nextID++
	p_user.Id = nextID
	ListUsers[p_user.Id] = p_user
}

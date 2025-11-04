package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nghSia/Mini-CRM/internal/storage"
)

// DisplayMenu affiche le menu principal de l'application
func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== Mini-CRM Menu ===")
		fmt.Println("1) : Ajouter un contact")
		fmt.Println("2) : Lister tous les contacts")
		fmt.Println("3) : Liste les info d'un contact")
		fmt.Println("4) : Mettre a jour un contact")
		fmt.Println("5) : Supprimer un contact")
		fmt.Println("6) Quitter l'application")

		choice := readInt(reader, "Entrez vos choix : ")

		switch choice {
		case 1:
			handleAddContact(reader, store)
		case 2:
			handleGetAllContact(store)
		case 3:
			handleGetContactByID(reader, store)
		case 4:
			handleUpdateContact(reader, store)
		case 5:
			handleDeleteContact(reader, store)
		case 6:
			fmt.Println("Fermeture...")
			return
		default:
			fmt.Println(choice)
		}
	}
}

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Entrez nom utilisateur : ")
	i_username, _ := reader.ReadString('\n')
	i_username = strings.TrimSpace(i_username)

	fmt.Print("Entrez le mail utilisateur : ")
	i_userMail, _ := reader.ReadString('\n')
	i_userMail = strings.TrimSpace(i_userMail)

	newUser := &storage.Contact{Name: i_username, Email: i_userMail}

	store.Add(newUser)
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Entrez l'Id d'utilisateur a supprimer : ")
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	if err != nil {
		fmt.Errorf("❌ Id invalide, veuillez entrer un nombre entier.")
		return
	}

	err = store.Delete(i_index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("✅ Contact supprimé avec succès")
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Entrez l'Id d'utilisateur a mettre a jour : ")
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	contactsLength, err := store.GetAll()

	if len(contactsLength) < i_index {
		fmt.Println("❌ Id invalide")
		return
	}

	fmt.Print("Entrez le nouveau nom : ")
	i_name, _ := reader.ReadString('\n')
	i_name = strings.TrimSpace(i_name)

	fmt.Print("Entrez le nouveau mail : ")
	i_email, _ := reader.ReadString('\n')
	i_email = strings.TrimSpace(i_email)

	err = store.Update(i_index, i_name, i_email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("✅ Contact mis à jour avec succès")
}

func handleGetAllContact(store storage.Storer) {
	contacts, err := store.GetAll()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", contact.Id, contact.Name, contact.Email)
	}
}

func handleGetContactByID(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Entrez l'Id d'utilisateur à afficher : ") // Correction du message
	i_indexStr, _ := reader.ReadString('\n')
	i_indexStr = strings.TrimSpace(i_indexStr)
	i_index, err := strconv.Atoi(i_indexStr)

	if err != nil {
		fmt.Println("❌ Id invalide, veuillez entrer un nombre entier.") // Changé Printf en Println
		return
	}

	contact, err := store.GetById(i_index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ID: %d | Nom: %s | Email: %s\n", contact.Id, contact.Name, contact.Email)
}

// ReadInt assure la valeur renseingé par l'utilisateur soit une valeur attendue
func readInt(reader *bufio.Reader, inputMessage string) int {
	for {
		fmt.Print(inputMessage)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			fmt.Println("❌ Veuillez entrer une valeur valide")
			continue
		}

		value, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("❌ Veuillez entrer un nombre entier.")
			continue
		}

		if value < 1 || value > 6 {
			fmt.Println("❌ Veuillez entrer un nombre entre 1 et 6.")
			continue
		}
		return value
	}
}

package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	crudcontact "github.com/nghSia/Mini-CRM/crudContact"
	"github.com/nghSia/Mini-CRM/user"
)

// DisplayMenu affiche le menu principal de l'application
func DisplayMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== Mini-CRM Menu ===")
		fmt.Println("1) : Ajouter un contact")
		fmt.Println("2) : Lister tous les contacts")
		fmt.Println("3) : Supprimer un contact")
		fmt.Println("4) : Mettre a jour un contact")
		fmt.Println("5) Quitter l'application")

		choice := readInt(reader, "Entrez vos choix : ")

		switch choice {
		case 1:
			fmt.Print("Entrez nom utilisateur : ")
			i_username, _ := reader.ReadString('\n')
			i_username = strings.TrimSpace(i_username)
			fmt.Print("Entrez le mail utilisateur : ")
			i_userMail, _ := reader.ReadString('\n')
			i_userMail = strings.TrimSpace(i_userMail)
			newUser := user.User{Name: i_username, Email: i_userMail}
			crudcontact.AddUser(newUser)
		case 2:
			crudcontact.GetUsers()
		case 3:
			fmt.Print("Entrez l'Id d'utilisateur a supprimer : ")
			i_indexStr, _ := reader.ReadString('\n')
			i_indexStr = strings.TrimSpace(i_indexStr)
			i_index, err := strconv.Atoi(i_indexStr)
			if err != nil {
				fmt.Println("❌ Id invalide, veuillez entrer un nombre entier.")
				continue
			}
			crudcontact.DeleteUser(i_index)
		case 4:
			fmt.Print("Entrez l'Id d'utilisateur a mettre a jour : ")
			i_indexStr, _ := reader.ReadString('\n')
			i_indexStr = strings.TrimSpace(i_indexStr)
			i_index, err := strconv.Atoi(i_indexStr)
			if err != nil {
				fmt.Println("❌ Id invalide, veuillez entrer un nombre entier.")
				continue
			}
			fmt.Print("Entrez le nouveau nom : ")
			i_name, _ := reader.ReadString('\n')
			i_name = strings.TrimSpace(i_name)
			fmt.Print("Entrez le nouveau mail : ")
			i_email, _ := reader.ReadString('\n')
			i_email = strings.TrimSpace(i_email)
			crudcontact.UpdateUser(i_index, user.User{Name: i_name, Email: i_email})
		case 5:
			fmt.Println("Fermeture...")
			return
		default:
			fmt.Println(choice)
		}
	}
}

// ReadInt assure la valeur renseingé par l'utilisateur soit une valeur attendue
func readInt(reader *bufio.Reader, inputMessage string) int {
	for {
		fmt.Print(inputMessage)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		value, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("❌ Veuillez entrer un nombre entier.")
			continue
		}

		if value < 1 || value > 5 {
			fmt.Println("❌ Veuillez entrer un nombre entre 1 et 5.")
			continue
		}
		return value
	}
}

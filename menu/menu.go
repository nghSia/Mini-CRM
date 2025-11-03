package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UserChoice uint8

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
		fmt.Println("choix : ", choice)
		if choice == 5 {
			fmt.Println("Fermeture...")
			return
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

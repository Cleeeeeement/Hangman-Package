package hangman

import (
	"fmt"
)

func GetInput(word string) string {
	var entrer string
	var end bool
	for !end {

		fmt.Print("Choose  : ")
		fmt.Scanln(&entrer)
		if len(entrer) == 0 {
			fmt.Println("Vous devez entrer au minimum UN caractère !")
		} else if len(entrer) == 1 {
			if rune(entrer[0]) >= 97 && rune(entrer[0]) <= 122 {
				return entrer
			} else if rune(entrer[0]) >= 65 && rune(entrer[0]) <= 90 {
				entrer = string(rune(entrer[0]) + 32)
				return entrer
			} else {
				fmt.Println("Vous ne pouvez entrer que des LETTRES !")
			}
		} else {
			return entrer
		}
	}
	return entrer
}

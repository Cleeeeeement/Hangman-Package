package hangman

import (
	"math/rand"
	"time"
)

func Randomletters(word string) []int {
	// calcul du nombre de lettre a afficher
	nblettre := len(word)/2 - 1
	// boucle de select lettre
	var tabposprise = make([]int, nblettre)
	min := 0
	max := len(word) - 1
	rand.Seed(time.Now().UnixNano())
	posrandom := (rand.Intn(max-min+1) + min)
	for i := 0; i < nblettre; i++ {
		posfound := false
		for posfound != true {
			rand.Seed(time.Now().UnixNano())
			posrandom = (rand.Intn(max-min+1) + min)
			if in_array(posrandom, tabposprise) {
				continue
			} else {
				posfound = true
				tabposprise = append(tabposprise, posrandom)
			}
		}
	}
	return tabposprise
}

func in_array(val int, array []int) bool {
	for i := range array {
		if array[i] == val {
			return true
		}
	}
	return false
}

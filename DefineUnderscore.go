package hangman

import (
	"math/rand"
	"time"
)

// bite//
func Randomletters(word string) []string {
	var underscore []string
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, "_")
	}
	for i := 0; i < len(word)/2-1; {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(len(word) - 1)
		if underscore[random] != string(word[random]) {
			underscore[random] = string(word[random])
			i++
		}
	}
	return underscore
}

package hangman

func GuessWord(word string, underscore []string, input string, lettre_list []string) []string {

	bad_word := []string{"B", "A", "D", "!"}
	same_word := []string{"S", "A", "M", "E", "!"}

	if len(input) == 1 {
		for _, l := range lettre_list {
			if l == input {
				return same_word
			}
		}
		for k, i := range word {
			if input == string(i) {
				underscore[k] = input
			}
		}
	} else if input == word {
		for k, m := range word {
			underscore[k] = string(m)
		}
	} else {
		return bad_word
	}
	return underscore
}

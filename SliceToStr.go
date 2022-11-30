package hangman

func SliceToStr(slice []string) string {
	var word string
	for _, i := range slice {
		word += i
	}
	return word
}

func StrtoSlice(str string) []string {
	var word []string
	for _, i := range str {
		word = append(word, string(i))
	}
	return word
}

package helper

import "strings"

func ToTitleCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}

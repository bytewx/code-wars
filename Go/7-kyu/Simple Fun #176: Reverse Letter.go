package kata

import (
	"unicode"
)

func ReverseLetters(s string) string {
	letters := []rune{}

	for _, char := range s {
		if unicode.IsLetter(char) {
			letters = append(letters, char)
		}
	}

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	return string(letters)
}

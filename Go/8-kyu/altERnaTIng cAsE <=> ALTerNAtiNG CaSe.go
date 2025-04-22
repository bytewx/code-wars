package kata

import (
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	var result strings.Builder

	for _, char := range str {
		if unicode.IsLower(char) {
			result.WriteRune(unicode.ToUpper(char))
		} else if unicode.IsUpper(char) {
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

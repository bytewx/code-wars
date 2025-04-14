package kata

import (
	"strings"
)

func bandNameGenerator(word string) string {
	if len(word) == 0 {
		return ""
	}
	first := word[0]
	last := word[len(word)-1]

	if first == last {
		capitalized := strings.Title(word)
		return capitalized + word[1:]
	}

	return "The " + strings.Title(word)
}

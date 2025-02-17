package kata

import (
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	freq := make(map[rune]int)
	
	for _, char := range str {
		freq[char]++
	}
	
	result := make([]string, len(str))
	for i, char := range str {
		result[i] = strconv.Itoa(freq[char])
	}
	
	return strings.Join(result, sep)
}

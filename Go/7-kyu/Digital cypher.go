package kata

import "strconv"

func Encode(str string, key int) []int {
	keyStr := strconv.Itoa(key)
	keyLen := len(keyStr)
	result := make([]int, len(str))

	for i, char := range str {
		charValue := int(char - 'a' + 1) 
		keyDigit := int(keyStr[i%keyLen] - '0')
		result[i] = charValue + keyDigit
	}

	return result
}

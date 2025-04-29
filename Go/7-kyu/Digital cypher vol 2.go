package kata

import (
	"strings"
  "strconv"
)

func Decode(code []int, key int) string {
	numToLetter := make(map[int]byte)
	for i := 1; i <= 26; i++ {
		numToLetter[i] = byte('a' + i - 1)
	}

	var decoded strings.Builder
	keyStr := []byte(strconv.Itoa(key))

	for i, num := range code {
		originalNum := num - int(keyStr[i%len(keyStr)] - '0')

		decoded.WriteByte(numToLetter[originalNum])
	}

	return decoded.String()
}

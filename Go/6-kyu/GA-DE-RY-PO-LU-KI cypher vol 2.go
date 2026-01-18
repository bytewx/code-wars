package kata

import "unicode"

func buildCipher(key string) map[rune]rune {
	cipher := make(map[rune]rune)

	for i := 0; i < len(key)-1; i += 2 {
		a := rune(key[i])
		b := rune(key[i+1])

		cipher[a] = b
		cipher[b] = a

		cipher[unicode.ToUpper(a)] = unicode.ToUpper(b)
		cipher[unicode.ToUpper(b)] = unicode.ToUpper(a)
	}

	return cipher
}

func transform(str, key string) string {
	cipher := buildCipher(key)
	result := make([]rune, 0, len(str))

	for _, ch := range str {
		if mapped, ok := cipher[ch]; ok {
			result = append(result, mapped)
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

func Encode(str, key string) string {
	return transform(str, key)
}

func Decode(str, key string) string {
	return transform(str, key)
}

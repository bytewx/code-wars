package kata

func GetTheVowels(word string) int {
	vowels := "aeiou"
	expected := 0
	count := 0

	for i := 0; i < len(word); i++ {
		if word[i] == vowels[expected] {
			count++
			expected = (expected + 1) % len(vowels)
		}
	}

	return count
}

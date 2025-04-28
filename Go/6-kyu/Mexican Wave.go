package kata

func wave(words string) []string {
	var result []string

	for i, ch := range words {
		if ch == ' ' {
			continue
		}

		wave := words[:i] + string(ch-32) + words[i+1:]
		result = append(result, wave)
	}

	return result
}

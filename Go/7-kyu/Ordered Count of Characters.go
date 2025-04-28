package orderedcount

func OrderedCount(text string) []Tuple {
    counts := make(map[rune]int)
    seen := []rune{}

    for _, ch := range text {
        if _, exists := counts[ch]; !exists {
            seen = append(seen, ch)
        }
        counts[ch]++
    }

    result := make([]Tuple, 0, len(seen))
    for _, ch := range seen {
        result = append(result, Tuple{Char: ch, Count: counts[ch]})
    }

    return result
}

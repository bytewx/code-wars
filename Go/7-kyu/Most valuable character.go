package kata

func Solve(s string) rune {
    firstIndex := make(map[rune]int)
    lastIndex := make(map[rune]int)

    for i, ch := range s {
        if _, exists := firstIndex[ch]; !exists {
            firstIndex[ch] = i
        }
        lastIndex[ch] = i
    }

    var result rune
    maxValue := -1

    for ch := 'a'; ch <= 'z'; ch++ {
        fi, ok := firstIndex[ch]
        if !ok {
            continue
        }
        li := lastIndex[ch]
        value := li - fi
        if value > maxValue {
            maxValue = value
            result = ch
        }
    }

    return result
}

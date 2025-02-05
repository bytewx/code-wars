package kata

import "unicode"

func solve(words []string) []int {
    results := make([]int, len(words))
    
    for i, word := range words {
        count := 0
        for j, char := range word {
            if unicode.ToLower(char)-'a' == int32(j) {
                count++
            }
        }
        results[i] = count
    }
    
    return results
}

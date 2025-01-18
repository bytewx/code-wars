package kata

import "strings"

func IsPalindrome(str string) bool {
    if len(str) <= 1 {
        return true
    }

    r := []rune(strings.ToLower(str))
    
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        if r[i] != r[j] {
            return false
        }
    }

    return true
}

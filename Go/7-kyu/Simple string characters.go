package kata

import "unicode"

func Solve(s string) []int {
    res := []int{0, 0, 0, 0}

    for _, ch := range s {
        switch {
        case unicode.IsUpper(ch):
            res[0]++
        case unicode.IsLower(ch):
            res[1]++
        case unicode.IsDigit(ch):
            res[2]++
        default:
            res[3]++
        }
    }

    return res
}

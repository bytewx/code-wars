package kata

import "strconv"

func FakeBin(x string) string {
    result := ""
    for _, char := range x {
        digit, _ := strconv.Atoi(string(char))
        if digit < 5 {
            result += "0"
        } else {
            result += "1"
        }
    }
    return result
}

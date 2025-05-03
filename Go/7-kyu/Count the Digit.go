package kata

import (
    "strconv"
    "strings"
)

func NbDig(n int, d int) int {
    count := 0
    digit := strconv.Itoa(d)

    for k := 0; k <= n; k++ {
        square := k * k
        strSquare := strconv.Itoa(square)
        count += strings.Count(strSquare, digit)
    }

    return count
}

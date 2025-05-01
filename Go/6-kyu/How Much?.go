package kata

import "fmt"

func HowMuch(m int, n int) [][3]string {
    if m > n {
        m, n = n, m
    }

    var res [][3]string

    for f := m; f <= n; f++ {
        if (f-1)%9 == 0 && (f-2)%7 == 0 {
            c := (f - 1) / 9
            b := (f - 2) / 7
            entry := [3]string{
                fmt.Sprintf("M: %d", f),
                fmt.Sprintf("B: %d", b),
                fmt.Sprintf("C: %d", c),
            }
            res = append(res, entry)
        }
    }

    return res
}

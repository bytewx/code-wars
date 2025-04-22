package kata

func ExpressionMatter(a int, b int, c int) int {
    res1 := a + b + c
    res2 := a * b * c
    res3 := a + (b * c)
    res4 := (a + b) * c
    res5 := a * (b + c)
    res6 := (a * b) + c

    max := res1
    if res2 > max {
        max = res2
    }
    if res3 > max {
        max = res3
    }
    if res4 > max {
        max = res4
    }
    if res5 > max {
        max = res5
    }
    if res6 > max {
        max = res6
    }

    return max
}

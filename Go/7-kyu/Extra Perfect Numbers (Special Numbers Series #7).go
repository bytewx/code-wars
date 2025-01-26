package kata

func ExtraPerfect(n int) []int {
    var result []int
    for i := 1; i <= n; i += 2 {
        result = append(result, i)
    }
    return result
}

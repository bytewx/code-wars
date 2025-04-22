package kata

func PowersOfTwo(n int) []uint64 {
    var result []uint64
    for i := 0; i <= n; i++ {
        result = append(result, 1 << i)
    }
    return result
}

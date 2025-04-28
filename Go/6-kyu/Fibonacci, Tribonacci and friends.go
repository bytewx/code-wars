package kata

func Xbonacci(signature []int, n int) []int {
    x := len(signature)

    if n == 0 {
        return []int{}
    }

    if n <= x {
        return signature[:n]
    }

    result := make([]int, n)
    copy(result, signature)

    for i := x; i < n; i++ {
        var sum int
        for j := i - x; j < i; j++ {
            sum += result[j]
        }
        result[i] = sum
    }

    return result
}

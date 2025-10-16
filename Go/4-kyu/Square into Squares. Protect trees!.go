package kata

func Decompose(n int64) []int64 {
    res := decomposeHelper(n * n, n-1)
    return res
}

func decomposeHelper(remaining, max int64) []int64 {
    if remaining == 0 {
        return []int64{}
    }
    for i := max; i > 0; i-- {
        sq := i * i
        if sq <= remaining {
            sub := decomposeHelper(remaining - sq, i-1)
            if sub != nil {
                return append(sub, i)
            }
        }
    }
    return nil
}

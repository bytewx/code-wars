package kata

func IsNegativeZero(n float64) bool {
    return n == 0 && (1/n) < 0
}

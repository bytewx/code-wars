package kata

func GetSum(a, b int) int {
    if a > b {
        a, b = b, a
    }
    
    sum := 0
    for i := a; i <= b; i++ {
        sum += i
    }
    return sum
}

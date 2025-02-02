package kata

func Seven(n int64) []int {
    steps := 0
  
    for n >= 100 {
        n = (n / 10) - 2*(n % 10)
        steps++
    }
  
    return []int{int(n), steps}
}

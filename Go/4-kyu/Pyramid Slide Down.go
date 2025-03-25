package kata

func LongestSlideDown(pyramid [][]int) int {
    for i := len(pyramid) - 2; i >= 0; i-- {
        for j := 0; j < len(pyramid[i]); j++ {
            pyramid[i][j] += max(pyramid[i+1][j], pyramid[i+1][j+1])
        }
    }
  
    return pyramid[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

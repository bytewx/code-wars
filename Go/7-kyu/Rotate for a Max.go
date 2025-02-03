package kata

import (
  "fmt"
  "strconv"
)

func MaxRot(n int64) int64 {
    s := []rune(fmt.Sprintf("%d", n))
    maxVal := n
    
    for i := 0; i < len(s)-1; i++ {
        rotated := append(s[:i], append(s[i+1:], s[i])...)
        newVal, _ := strconv.ParseInt(string(rotated), 10, 64)
        if newVal > maxVal {
            maxVal = newVal
        }
        s = rotated
    }
    
    return maxVal
}

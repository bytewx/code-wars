package kata

import "fmt"

func countSheep(num int) string {
    result := ""
    for i := 1; i <= num; i++ {
        result += fmt.Sprintf("%d sheep...", i)
    }
    return result
}

package kata

import (
  "fmt"
  "strings"
)

func MultiTable(number int) string {
  var result strings.Builder
  
  for i := 1; i <= 10; i++ {
    result.WriteString(fmt.Sprintf("%d * %d = %d", i, number, i * number))
    
    if i < 10 {
      result.WriteString("\n")
    }
  }
  
  return result.String()
}

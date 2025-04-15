package kata

import (
  "sort"
  "strings"
)

func TwoSort(arr []string) string {
  sort.Strings(arr)
  first := arr[0]  

  var sb strings.Builder
  for i, ch := range first {
    sb.WriteString(string(ch))
    if i < len(first)-1 {
      sb.WriteString("***") 
    }
  }

  return sb.String()
}

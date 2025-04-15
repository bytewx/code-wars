package kata

import "strings"

func Contamination(text, char string) string {
  var result []string
  
  for i := 1; i <= len(text); i++ {
    result = append(result, char)
  }
  
  return strings.Join(result, "")
}

package kata

import "strings"

func Vaporcode(s string) string {
  var result strings.Builder
  
  processedString := strings.ReplaceAll(s, " ", "")
  
  for i := 0; i < len(processedString); i++ {
    result.WriteString(string(processedString[i]))
    
    if i < len(processedString) - 1 {
      result.WriteString("  ")
    }
  }
  
  return strings.ToUpper(result.String())
}

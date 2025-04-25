package kata

import "strings"

func duplicate_count(s1 string) int {
  s1 = strings.ToLower(s1)
  
  countMap := make(map[rune]int)
  
  for _, char := range s1 {
    countMap[char]++
  }
  
  duplicateCount := 0
  
  for _, count := range countMap {
    if count > 1 {
      duplicateCount++
    }
  }
  
  return duplicateCount
}

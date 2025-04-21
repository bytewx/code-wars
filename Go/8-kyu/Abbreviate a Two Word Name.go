package kata

import (
  "fmt"
  "strings"
)

func AbbrevName(name string) string {
  nameSlice := strings.Split(name, " ")
  
  firstLetter, secondLetter := string(nameSlice[0][0]), string(nameSlice[1][0])
  
  return fmt.Sprintf("%s.%s", strings.ToUpper(firstLetter), strings.ToUpper(secondLetter))
}

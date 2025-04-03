package kata

import (
  "strings"
  "strconv"
)

func ToCsvText(array [][]int) string {
  var rows []string
  
  for _, row := range array {
    var rowString []string
    
    for _, number := range row {
      rowString = append(rowString, strconv.Itoa(number))
    }
    
    rows = append(rows, strings.Join(rowString, ","))
  }
  
  return strings.Join(rows, "\n")
}

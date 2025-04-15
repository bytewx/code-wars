package kata

import "strconv"

func Digitize(n int) []int {
  var result []int
  
  numberString := strconv.Itoa(n)
  
  for i := len(numberString) - 1; i >= 0; i-- {
    digit, _ := strconv.Atoi(string(numberString[i]))
    result = append(result, digit)
  }
  
  return result
}

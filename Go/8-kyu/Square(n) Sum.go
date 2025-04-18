package kata

import "math"

func SquareSum(numbers []int) int {
  var sum float64
  
  for _, value := range numbers {
    sum = sum + math.Pow(float64(value), 2)
  }
  
  return int(sum)
}

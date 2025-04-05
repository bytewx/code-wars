package kata

import "math"

func SumCubes(n int) int {
  var sum int
  
  for i := 1; i <= n; i++ {
    sum = sum + int(math.Pow(float64(i), 3))
  }
  
  return sum
}

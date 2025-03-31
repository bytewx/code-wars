package kata

import "math"

func century(year int) int {
  if year % 100 == 0 {
    return int(math.Floor(float64(year) / 100))
  }
  
  return int(math.Floor(float64(year) / 100)) + 1
}

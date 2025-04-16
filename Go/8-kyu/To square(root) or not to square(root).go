package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
    var result []int

    for _, value := range arr {
        sqrt := math.Sqrt(float64(value))
      
        if sqrt == math.Floor(sqrt) {
            result = append(result, int(sqrt))
        } else {
            result = append(result, value * value)
        }
    }

    return result
}

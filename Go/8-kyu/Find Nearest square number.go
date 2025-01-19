package kata

import (
	"math"
)

func NearestSq(n int) int {
	sqrt := math.Sqrt(float64(n))
	nearestInt := int(math.Round(sqrt))

	return nearestInt * nearestInt
}

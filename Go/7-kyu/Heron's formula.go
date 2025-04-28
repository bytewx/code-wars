package kata

import (
	"math"
)

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	area = math.Round(area*100) / 100
	return area
}

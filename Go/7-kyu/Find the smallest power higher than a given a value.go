package kata

import "math"

func FindNextPower(val, pow int) int {
	fval := float64(val)
	fexp := float64(pow)

	base := math.Pow(fval, 1.0/fexp)
	nextBase := math.Ceil(base)

	if math.Pow(nextBase, fexp) <= fval {
		nextBase++
	}

	return int(math.Pow(nextBase, fexp))
}

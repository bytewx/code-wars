package kata

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solve(s int, g int) []int {
	if s%g != 0 {
		return []int{-1, -1}
	}

	k := s / g

	for x := 1; x <= k/2; x++ {
		y := k - x
		if gcd(x, y) == 1 {
			return []int{g * x, g * y}
		}
	}

	return []int{-1, -1}
}

package kata

func Solequa(n int) [][]int {
	result := [][]int{}

	for a := 1; a*a <= n; a++ {
		if n%a != 0 {
			continue
		}
		b := n / a

		if (a + b)%2 != 0 || (b - a)%4 != 0 {
			continue
		}

		x := (a + b) / 2
		y := (b - a) / 4

		if x >= 0 && y >= 0 {
			result = append(result, []int{x, y})
		}
	}

	return result
}

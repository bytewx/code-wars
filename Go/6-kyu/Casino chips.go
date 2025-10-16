package kata

func Solve(arr []int) int {
	if len(arr) != 3 {
		return 0
	}
	a, b, c := arr[0], arr[1], arr[2]
	sum := a + b + c

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	d1 := sum / 2
	d2 := sum - max
	if d1 < d2 {
		return d1
	}
	return d2
}

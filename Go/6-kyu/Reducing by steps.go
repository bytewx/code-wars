package kata

func Gcdi(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func Som(x, y int) int {
	return x + y
}

func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Lcmu(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	xAbs, yAbs := x, y
	if xAbs < 0 {
		xAbs = -xAbs
	}
	if yAbs < 0 {
		yAbs = -yAbs
	}
	return xAbs / Gcdi(xAbs, yAbs) * yAbs
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	res := make([]int, len(arr))
	acc := init

	for i, v := range arr {
		acc = f(acc, v)
		res[i] = acc
	}
	return res
}

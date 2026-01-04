package kata

import "strconv"

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SumFracts(arr [][]int) string {
	if len(arr) == 0 {
		return "0"
	}

	var num int64 = 0
	var den int64 = 1

	for _, f := range arr {
		a := int64(f[0])
		b := int64(f[1])

		num = num*b + a*den
		den = den * b

		g := gcd(num, den)
		num /= g
		den /= g
	}

	if num%den == 0 {
		return strconv.FormatInt(num/den, 10)
	}

	return strconv.FormatInt(num, 10) + "/" + strconv.FormatInt(den, 10)
}

package kata

import (
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func Strong(n int) string {
	sum := 0
	nStr := strconv.Itoa(n)
	for _, digit := range nStr {
		num := int(digit - '0')
		sum += factorial(num)
	}

	if sum == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

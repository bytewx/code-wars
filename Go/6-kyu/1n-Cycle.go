package kata

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Cycle(n int) int {
	if gcd(n, 10) != 1 {
		return -1
	}

	remainder := 1
	length := 1

	for (remainder*10)%n != 1 {
		remainder = (remainder * 10) % n
		length++
	}

	return length
}

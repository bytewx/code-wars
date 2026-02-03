package kata

func ComputeDepth(n uint) uint {
	seen := make(map[rune]bool)
	var depth uint = 0

	for len(seen) < 10 {
		depth++
		value := n * depth

		for _, digit := range []rune(fmtUint(value)) {
			seen[digit] = true
		}
	}

	return depth
}

func fmtUint(n uint) string {
	if n == 0 {
		return "0"
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{'0' + byte(n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}

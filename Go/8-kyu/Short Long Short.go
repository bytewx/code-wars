package kata

func Solution(a, b string) string {
	if len(a) < len(b) {
		return a + b + a
	}
	return b + a + b
}

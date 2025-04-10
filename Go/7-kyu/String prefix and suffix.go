package kata

func Solve(s string) int {
	maxLen := 0
	n := len(s)
	for i := 1; i < n; i++ {
		if s[:i] == s[n-i:] {
			maxLen = i
		}
	}
	return maxLen
}

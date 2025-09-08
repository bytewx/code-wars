package kata

func ExpSum(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	dp := make([]uint64, n+1)
	dp[0] = 1

	for k := uint64(1); k <= n; k++ {
		for i := k; i <= n; i++ {
			dp[i] += dp[i-k]
		}
	}
	return dp[n]
}

package kata

func Diagonal(n, p int) int {
	return binomial(n+1, p+1)
}

func binomial(n, k int) int {
	if k > n-k {
		k = n - k
	}

	result := 1
	for i := 1; i <= k; i++ {
		result = result * (n - k + i) / i
	}
	return result
}

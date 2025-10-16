package kata

import (
	"fmt"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPairwiseCoprime(sys []int64) bool {
	for i := 0; i < len(sys); i++ {
		for j := i + 1; j < len(sys); j++ {
			if gcd(sys[i], sys[j]) != 1 {
				return false
			}
		}
	}
	return true
}

func FromNbToStr(n int64, sys []int64) string {
	if !isPairwiseCoprime(sys) {
		return "Not applicable"
	}

	product := int64(1)
	for _, m := range sys {
		product *= m
	}

	if product <= n {
		return "Not applicable"
	}

	result := ""
	for _, m := range sys {
		residue := n % m
		result += fmt.Sprintf("-%d-", residue)
	}

	return result
}

package kata

import (
	"math/big"
	"sort"
)

func sumDigits(x *big.Int) int {
	s := x.String()
	sum := 0
	for _, ch := range s {
		sum += int(ch - '0')
	}
	return sum
}

func PowerSumDigTerm(n int) *big.Int {
	results := []*big.Int{}

	for a := 2; a < 200; a++ {
		base := big.NewInt(int64(a))
		pow := new(big.Int).Set(base)

		for b := 2; b < 15; b++ {
			pow = new(big.Int).Mul(pow, base)

			if pow.Cmp(big.NewInt(10)) < 0 {
				continue
			}

			if sumDigits(pow) == a {
				results = append(results, new(big.Int).Set(pow))
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Cmp(results[j]) < 0
	})

	return results[n-1]
}

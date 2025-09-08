package kata

import (
	"math/big"
)

func EasyLine(n int) string {
	if n == 0 {
		return "1"
	}

	res := new(big.Int)
	res.Binomial(int64(2*n), int64(n))

	return res.String()
}

package kata

func PyramidHeight(n int) int {
	if n <= 0 {
		return 0
	}
	N := int64(n)

	sumSquares := func(h int64) int64 {
		return h * (h + 1) * (2*h + 1) / 6
	}

	lo, hi := int64(0), int64(1)
	for sumSquares(hi) <= N {
		hi *= 2
	}

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if sumSquares(mid) <= N {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return int(lo)
}

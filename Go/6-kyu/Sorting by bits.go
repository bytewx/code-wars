package kata

import (
	"math/bits"
	"sort"
)

func countOnBits(n int) int {
	return bits.OnesCount(uint(n))
}

func SortByBit(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		bitsI := countOnBits(arr[i])
		bitsJ := countOnBits(arr[j])
		if bitsI == bitsJ {
			return arr[i] < arr[j]
		}
		return bitsI < bitsJ
	})
}

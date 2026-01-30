package kata

import (
	"strconv"
)

func Smallest(n int64) []int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)

	bestVal := n
	bestI, bestJ := 0, 0

	move := func(s string, i, j int) string {
		d := s[i]
		rem := make([]byte, 0, len(s)-1)
		rem = append(rem, s[:i]...)
		rem = append(rem, s[i+1:]...)

		out := make([]byte, 0, len(s))
		out = append(out, rem[:j]...)
		out = append(out, d)
		out = append(out, rem[j:]...)
		return string(out)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			candStr := move(s, i, j)
			candVal, _ := strconv.ParseInt(candStr, 10, 64) 

			if candVal < bestVal ||
				(candVal == bestVal && (i < bestI || (i == bestI && j < bestJ))) {
				bestVal = candVal
				bestI = i
				bestJ = j
			}
		}
	}

	return []int64{bestVal, int64(bestI), int64(bestJ)}
}

package kata

import (
	"sort"
)

func Nico(key, message string) string {
	k := len(key)

	type pair struct {
		ch  rune
		idx int
	}

	pairs := make([]pair, k)
	for i, ch := range key {
		pairs[i] = pair{ch, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].ch < pairs[j].ch
	})

	numKey := make([]int, k)
	for i, p := range pairs {
		numKey[p.idx] = i
	}

	for len(message)%k != 0 {
		message += " "
	}

	result := make([]rune, 0, len(message))
	rows := len(message) / k

	for r := 0; r < rows; r++ {
		row := []rune(message[r*k : (r+1)*k])
		sortedRow := make([]rune, k)
		for i := 0; i < k; i++ {
			sortedRow[numKey[i]] = row[i]
		}
		result = append(result, sortedRow...)
	}

	return string(result)
}

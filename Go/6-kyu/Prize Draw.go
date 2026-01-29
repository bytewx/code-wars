package kata

import (
	"sort"
	"strings"
	"unicode"
)

type participant struct {
	name  string
	score int
}

func nameSom(s string) int {
	sum := len(s)
	for _, r := range s {
		if unicode.IsLetter(r) {
			upper := unicode.ToUpper(r)
			sum += int(upper-'A') + 1
		}
	}
	return sum
}

func NthRank(st string, we []int, n int) string {
	st = strings.TrimSpace(st)
	if st == "" {
		return "No participants"
	}

	names := strings.Split(st, ",")
	if n > len(names) {
		return "Not enough participants"
	}

	parts := make([]participant, 0, len(names))
	for i, name := range names {
		som := nameSom(name)
		score := som * we[i]
		parts = append(parts, participant{name: name, score: score})
	}

	sort.Slice(parts, func(i, j int) bool {
		if parts[i].score != parts[j].score {
			return parts[i].score > parts[j].score
		}
		return parts[i].name < parts[j].name
	})

	return parts[n-1].name
}

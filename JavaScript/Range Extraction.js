package kata

import (
	"strconv"
	"strings"
)

func Solution(list []int) string {
	if len(list) == 0 {
		return ""
	}

	var parts []string

	addRange := func(a, b int) {
		switch diff := b - a; {
		case diff >= 2: 
			parts = append(parts, strconv.Itoa(a)+"-"+strconv.Itoa(b))
		case diff == 1: 
			parts = append(parts, strconv.Itoa(a), strconv.Itoa(b))
		default: 
			parts = append(parts, strconv.Itoa(a))
		}
	}

	start, prev := list[0], list[0]
	for i := 1; i < len(list); i++ {
		x := list[i]
		if x == prev+1 {
			prev = x
			continue
		}

		addRange(start, prev)

		start, prev = x, x
	}

	addRange(start, prev)

	return strings.Join(parts, ",")
}

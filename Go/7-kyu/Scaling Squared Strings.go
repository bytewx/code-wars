package kata

import (
	"strings"
)

func Scale(s string, k, v int) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\n")
	var scaledLines []string

	for _, line := range lines {
		var hScaledLine strings.Builder
		for _, ch := range line {
			hScaledLine.WriteString(strings.Repeat(string(ch), k))
		}
		
		scaledSegment := strings.Repeat(hScaledLine.String()+"\n", v)
		scaledLines = append(scaledLines, strings.TrimRight(scaledSegment, "\n"))
	}

	return strings.Join(scaledLines, "\n")
}

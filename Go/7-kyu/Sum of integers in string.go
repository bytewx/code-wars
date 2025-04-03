package kata

import (
	"regexp"
	"strconv"
)

func SumOfIntegersInString(s string) int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)

	sum := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		sum += num
	}

	return sum
}

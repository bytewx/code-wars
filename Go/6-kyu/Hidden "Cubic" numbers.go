package kata

import (
	"fmt"
	"strconv"
	"unicode"
)

func IsSumOfCubes(s string) string {
	var cubicNumbers []int
	sum := 0
	i := 0
	n := len(s)

	for i < n {
		if !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}

		j := i
		for j < n && unicode.IsDigit(rune(s[j])) && j-i < 3 {
			j++
		}

		numStr := s[i:j]
		num, _ := strconv.Atoi(numStr)

		if isCubic(num) {
			cubicNumbers = append(cubicNumbers, num)
			sum += num
		}

		i = j
	}

	if len(cubicNumbers) == 0 {
		return "Unlucky"
	}

	result := ""
	for _, num := range cubicNumbers {
		result += fmt.Sprintf("%d ", num)
	}
	result += fmt.Sprintf("%d Lucky", sum)
	return result
}

func isCubic(num int) bool {
	sum := 0
	temp := num
	for temp > 0 {
		d := temp % 10
		sum += d * d * d
		temp /= 10
	}
	return sum == num
}

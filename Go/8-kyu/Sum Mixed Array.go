package kata

import (
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0
	for _, v := range arr {
		switch value := v.(type) {
		case int:
			sum += value
		case string:
			num, err := strconv.Atoi(value)
			if err == nil {
				sum += num
			}
		}
	}
	return sum
}

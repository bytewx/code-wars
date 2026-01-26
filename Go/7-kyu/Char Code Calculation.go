package kata

import "strconv"

func Calc(s string) int {
	total1 := ""
	for _, ch := range s {
		total1 += strconv.Itoa(int(ch))
	}

	sum1 := 0
	sum2 := 0

	for _, d := range total1 {
		digit := int(d - '0')
		sum1 += digit

		if d == '7' {
			sum2 += 1
		} else {
			sum2 += digit
		}
	}

	return sum1 - sum2
}

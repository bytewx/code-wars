package kata

import (
	"time"
)

func UnluckyDays(year int) int {
	count := 0
	for month := time.January; month <= time.December; month++ {
		if time.Date(year, month, 13, 0, 0, 0, 0, time.UTC).Weekday() == time.Friday {
			count++
		}
	}
	return count
}

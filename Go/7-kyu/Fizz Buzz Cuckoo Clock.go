package kata

import "strconv"

func FizzBuzzCuckooClock(time string) string {
	hour, minute := time[:2], time[3:]
	h, _ := strconv.Atoi(hour)
	m, _ := strconv.Atoi(minute)
	
	if m == 0 {
		cuckooCount := h % 12
		if cuckooCount == 0 {
			cuckooCount = 12
		}
		cuckoo := ""
		for i := 0; i < cuckooCount; i++ {
			cuckoo += "Cuckoo "
		}
		return cuckoo[:len(cuckoo)-1]
	} else if m == 30 {
		// On the half hour
		return "Cuckoo"
	}
	
	result := ""
	if m % 3 == 0 {
		result += "Fizz"
	}
	if m % 5 == 0 {
		if result != "" {
			result += " "
		}
		result += "Buzz"
	}
	
	if result == "" {
		return "tick"
	}
	
	return result
}

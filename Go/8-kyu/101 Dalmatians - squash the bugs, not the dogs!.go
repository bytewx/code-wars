package kata

func HowManyDalmatians(number int) string {
	var result string

	if number == 101 {
		result = "101 DALMATIONS!!!"
	} else if number <= 10 {
		result = "Hardly any"
	} else if number <= 50 {
		result = "More than a handful!"
	} else if number <= 100 {
		result = "Woah that's a lot of dogs!"
	}
	return result
}

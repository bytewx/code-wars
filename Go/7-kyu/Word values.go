package kata

func NameValue(my_list []string) []int {
	result := make([]int, len(my_list))

	for i, s := range my_list {
		sum := 0
		for _, ch := range s {
			if ch >= 'a' && ch <= 'z' {
				sum += int(ch - 'a' + 1)
			}
		}
    
		result[i] = sum * (i + 1)
	}

	return result
}

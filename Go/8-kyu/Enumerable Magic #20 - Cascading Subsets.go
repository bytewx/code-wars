package kata

func EachCons(arr []int, n int) [][]int {
	if n <= 0 || n > len(arr) {
		return [][]int{}
	}

	result := [][]int{}
  
	for i := 0; i <= len(arr) - n; i++ {
		subset := make([]int, n)
		copy(subset, arr[i:i + n])
		result = append(result, subset)
	}
  
	return result
}

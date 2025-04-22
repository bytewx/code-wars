package kata

import "sort"

func MergeArrays(arr1, arr2 []int) []int {
	unique := make(map[int]bool)

	for _, val := range arr1 {
		unique[val] = true
	}
	for _, val := range arr2 {
		unique[val] = true
	}

	result := make([]int, 0, len(unique))
	for key := range unique {
		result = append(result, key)
	}

	sort.Ints(result)
	return result
}

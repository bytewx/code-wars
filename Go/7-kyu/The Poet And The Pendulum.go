package kata

import (
	"sort"
)

func Pendulum(values []int) []int {
	sort.Ints(values) 

	left := []int{}
	right := []int{}

	for i, v := range values {
		if i%2 == 0 {
			left = append([]int{v}, left...)
		} else {
			right = append(right, v) 
		}
	}

	return append(left, right...)
}

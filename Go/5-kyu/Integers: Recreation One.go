package kata

import "math"

func getDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

func ListSquared(m, n int) [][]int {
	result := [][]int{}
	for i := m; i <= n; i++ {
		divisors := getDivisors(i)
		sum := 0
		for _, d := range divisors {
			sum += d * d
		}
		sqrt := int(math.Sqrt(float64(sum)))
		if sqrt*sqrt == sum {
			result = append(result, []int{i, sum})
		}
	}
	return result
}

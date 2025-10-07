package kata

import "math"

func QueueTime(customers []int, n int) int {
    if len(customers) == 0 {
        return 0
    }

    tills := make([]int, n)

    for _, customer := range customers {
        minIndex := 0
        minTime := tills[0]
        for i := 1; i < n; i++ {
            if tills[i] < minTime {
                minTime = tills[i]
                minIndex = i
            }
        }

        tills[minIndex] += customer
    }

    total := 0
    for _, t := range tills {
        total = int(math.Max(float64(total), float64(t)))
    }

    return total
}

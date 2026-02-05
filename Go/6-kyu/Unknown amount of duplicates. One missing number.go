package kata

import "sort"

func FindDupsMiss(arr []int) (int, []int) {
    seen := make(map[int]int)

    minVal, maxVal := arr[0], arr[0]

    for _, v := range arr {
        seen[v]++
        if v < minVal {
            minVal = v
        }
        if v > maxVal {
            maxVal = v
        }
    }

    var dups []int
    missing := 0

    for i := minVal; i <= maxVal; i++ {
        if seen[i] == 0 {
            missing = i
        } else if seen[i] > 1 {
            dups = append(dups, i)
        }
    }

    sort.Ints(dups)
    return missing, dups
}

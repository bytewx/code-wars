package kata

func Dominator(a []int) int {
    freq := make(map[int]int)
    n := len(a)

    for _, num := range a {
        freq[num]++
      
        if freq[num] > n/2 {
            return num
        }
    }

    return -1
}

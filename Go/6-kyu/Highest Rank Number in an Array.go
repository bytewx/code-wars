package kata

func HighestRank(nums []int) int {
    frequency := make(map[int]int)
    maxFreq := 0
    maxNum := 0
    
    for _, num := range nums {
        frequency[num]++
        if frequency[num] > maxFreq || (frequency[num] == maxFreq && num > maxNum) {
            maxFreq = frequency[num]
            maxNum = num
        }
    }
    
    return maxNum
}

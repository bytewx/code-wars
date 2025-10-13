package kata

func MaximumSubarraySum(numbers []int) int {
    maxSoFar := 0
    currentSum := 0

    for _, n := range numbers {
        currentSum += n
        if currentSum < 0 {
            currentSum = 0
        }
        if currentSum > maxSoFar {
            maxSoFar = currentSum
        }
    }

    return maxSoFar
}

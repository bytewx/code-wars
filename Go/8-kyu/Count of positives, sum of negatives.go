package kata

func CountPositivesSumNegatives(numbers []int) []int {
  var positiveCount int
  var negativeSum int
  
  for _, value := range numbers {
    if value > 0 {
      positiveCount += 1
    } else {
      negativeSum += value
    }
  }
  
  return []int{positiveCount, negativeSum}
}

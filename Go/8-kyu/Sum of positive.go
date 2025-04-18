package kata

func PositiveSum(numbers []int) int {
  var sum int
  
  for _, value := range numbers {
    if value > 0 {
      sum = sum + value
    }
  }
  
  return sum
}

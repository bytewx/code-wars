package kata

func CountSheeps(numbers []bool) int {
  var sheeps int
  
  for _, value := range numbers {
    if value {
      sheeps++
    }
  }
  
  return sheeps
}

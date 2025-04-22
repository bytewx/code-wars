package kata

func Invert(arr []int) []int {
  var result []int
  
  for _, value := range arr {
    result = append(result, -value)
  }
  
  return result
}

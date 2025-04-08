package kata

func ReverseList(lst []int) []int {
  var result []int
  
  for i := len(lst) - 1; i >= 0; i-- {
    result = append(result, lst[i])
  }
  
  return result
}

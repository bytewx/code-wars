package kata

func ReverseSeq(n int) []int {
  var result []int
  
  for i := n; i > 0; i-- {
    result = append(result, i)
  }
  
  return result
}

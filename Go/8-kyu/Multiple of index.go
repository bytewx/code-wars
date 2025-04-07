package kata

func multipleOfIndex (ints []int) []int {
  var result []int
  
  for index, value := range ints {
    if index == 0 {
      continue
    }
    
    if value % index == 0 {
      result = append(result, value)
    }
  }
  
  return result
}

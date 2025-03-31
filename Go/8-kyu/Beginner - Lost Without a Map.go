package kata

func Maps(x []int) []int {
  var numbersSlice []int
  
  for index := range x {
    numbersSlice = append(numbersSlice, x[index] * 2)
  }
  
  return numbersSlice
}

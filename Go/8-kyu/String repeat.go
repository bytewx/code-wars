package kata

func RepeatStr(repetitions int, value string) string {
  var result string
  
  for i := 1; i <= repetitions; i++ {
    result += value
  }
  
  return result
}

package kata

func SimpleMultiplication(n int) int {
  if n % 2 == 0 {
    return n << 3
  }
  
  return (n << 3) + n
}

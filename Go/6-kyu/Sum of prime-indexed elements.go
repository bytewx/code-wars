package kata

func isPrime(n int) bool {
  if n < 2 {
    return false
  }
  
  for i := 2; i * i <= n; i++ {
    if n % i == 0 {
      return false
    }
  }
  
  return true
}

func Solve(arr []int) int {
  sum := 0

  for i := 0; i < len(arr); i++ {
    if isPrime(i) {
      sum += arr[i]
    }
  }

  return sum
}

package kata

import "fmt"

func Derive(coefficient, exponent int) string {
  return fmt.Sprintf("%vx^%v", coefficient * exponent, exponent - 1)
}

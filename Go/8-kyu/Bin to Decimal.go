package kata

import "strconv"

func BinToDec(bin string) int {
  result, _ := strconv.ParseInt(bin, 2, 64)
  
  return int(result)
}

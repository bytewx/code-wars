package kata

import "strings"

func ModifyMultiply(str string, loc, num int) string {
  var locationStringsSlice []string
  
  strSlice := strings.Split(str, " ")
  
  if num != 0 {
    for i := 1; i <= num; i++ {
      locationStringsSlice = append(locationStringsSlice, strSlice[loc])    
    }
  } else {
    locationStringsSlice = append(locationStringsSlice, strSlice[loc])   
  }
  
  return strings.Join(locationStringsSlice, "-")
}

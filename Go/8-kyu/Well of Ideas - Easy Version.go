package kata

func Well(x []string) string {
  goodCounter := 0
  
  for _, idea := range x {
    if idea == "good" {
      goodCounter++
    }
  }
  
  if goodCounter == 0 {
    return "Fail!"
  } else if goodCounter <= 2 {
    return "Publish!"
  }
  
  return "I smell a series!"
}

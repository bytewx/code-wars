package kata

func GetCount(str string) (count int) {  
  for _, character := range str {
    switch character {
      case 'a', 'e', 'i', 'o', 'u':
        count += 1
    }
  }
  
  return
}

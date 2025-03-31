package kata

func GetSize(w,h,d int) [2]int {
  surfaceArea := (2 * w * d) + (2 * h * d) + (2 * h * w)
  volume := w * h * d
  
  return [2]int{surfaceArea, volume}
}

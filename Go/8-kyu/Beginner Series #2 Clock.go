package kata

func Past(h, m, s int) int {
  hoursInMilliseconds := h * 3600000
  minutesInMilliseconds := m * 60000
  secondsInMilliseconds := s * 1000
  
  return hoursInMilliseconds + minutesInMilliseconds + secondsInMilliseconds
}

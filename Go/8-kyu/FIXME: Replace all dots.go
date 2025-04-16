package kata

import (
  "regexp" 
)

func ReplaceDots(str string) string {
  return regexp.MustCompile(`[.]`).ReplaceAllString(str, `-`)
}


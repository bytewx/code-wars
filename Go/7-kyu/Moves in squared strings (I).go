package kata

import (
	"strings"
)

func VertMirror(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = reverseString(line)
	}
	return strings.Join(lines, "\n")
}

func HorMirror(s string) string {
	lines := strings.Split(s, "\n")
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return strings.Join(lines, "\n")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Oper(f FParam, x string) string {
	return f(x)
}

type FParam func(string) string

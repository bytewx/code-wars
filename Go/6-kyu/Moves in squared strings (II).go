package kata

import "strings"

func Rot(s string) string {
    lines := strings.Split(s, "\n")
    n := len(lines)

    for i := 0; i < n/2; i++ {
        lines[i], lines[n-1-i] = reverse(lines[n-1-i]), reverse(lines[i])
    }

    if n%2 == 1 {
        lines[n/2] = reverse(lines[n/2])
    }

    return strings.Join(lines, "\n")
}

func SelfieAndRot(s string) string {
    lines := strings.Split(s, "\n")
    dots := strings.Repeat(".", len(lines[0]))

    var result []string

    for _, line := range lines {
        result = append(result, line+dots)
    }

    rotated := strings.Split(Rot(s), "\n")
    for _, line := range rotated {
        result = append(result, dots+line)
    }

    return strings.Join(result, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
    return f(x)
}

func reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

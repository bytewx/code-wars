package kata

func SortMyString(s string) string {
    var even, odd []rune

    for i, ch := range s {
        if i%2 == 0 {
            even = append(even, ch)
        } else {
            odd = append(odd, ch)
        }
    }

    return string(even) + " " + string(odd)
}

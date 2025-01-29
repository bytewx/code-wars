package kata

func WordsToMarks(s string) int {
    sum := 0
    for _, char := range s {
        sum += int(char - 'a' + 1)
    }
    return sum
}

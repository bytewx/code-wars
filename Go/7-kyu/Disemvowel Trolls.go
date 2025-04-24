package kata

func Disemvowel(comment string) string {
    vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
    result := ""
    
    for _, char := range comment {
        if !vowels[char] {
            result += string(char)
        }
    }
    
    return result
}

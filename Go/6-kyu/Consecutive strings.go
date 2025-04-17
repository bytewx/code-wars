package kata

func LongestConsec(strarr []string, k int) string {
    n := len(strarr)
    if n == 0 || k > n || k <= 0 {
        return ""
    }

    maxLen := 0
    longest := ""

    for i := 0; i <= n-k; i++ {
        temp := ""
        for j := 0; j < k; j++ {
            temp += strarr[i+j]
        }

        if len(temp) > maxLen {
            maxLen = len(temp)
            longest = temp
        }
    }

    return longest
}

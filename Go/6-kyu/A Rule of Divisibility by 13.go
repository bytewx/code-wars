package kata

func Thirt(n int) int {
    sequence := []int{1, 10, 9, 12, 3, 4}
    prev := -1

    for n != prev {
        prev = n
        sum := 0
        i := 0
        temp := n

        for temp > 0 {
            digit := temp % 10
            sum += digit * sequence[i%len(sequence)]
            temp /= 10
            i++
        }

        n = sum
    }

    return n
}

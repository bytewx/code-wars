package kata

func SequenceSum(begin, end, step int) int {
    if begin > end {
        return 0
    }
    
    sum := 0
    for i := begin; i <= end; i += step {
        if (i - begin) % step == 0 {
            sum += i
        }
    }
    
    return sum
}

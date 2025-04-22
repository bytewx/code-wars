package kata

func CalculateYears(humanYears int) (result [3]int) {
    result[0] = humanYears

    if humanYears == 1 {
        result[1] = 15
    } else if humanYears == 2 {
        result[1] = 15 + 9
    } else {
        result[1] = 15 + 9 + (humanYears-2)*4
    }

    if humanYears == 1 {
        result[2] = 15
    } else if humanYears == 2 {
        result[2] = 15 + 9
    } else {
        result[2] = 15 + 9 + (humanYears-2)*5
    }

    return result
}

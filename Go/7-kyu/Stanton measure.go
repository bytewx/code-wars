package kata

func StantonMeasure(arr []int) int {
    countOne := 0
    for _, value := range arr {
        if value == 1 {
            countOne++
        }
    }
    
    countN := 0
    for _, value := range arr {
        if value == countOne {
            countN++
        }
    }
    
    return countN
}

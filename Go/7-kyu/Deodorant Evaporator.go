package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {
    days := 0
    thresholdAmount := content * float64(threshold) / 100.0
    dailyLoss := 1.0 - float64(evapPerDay)/100.0

    for content > thresholdAmount {
        content *= dailyLoss
        days++
    }

    return days
}

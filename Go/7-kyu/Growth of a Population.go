package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
    years := 0
    population := p0
    
    for population < p {
        population += int(float64(population)*percent/100) + aug
        years++
    }
    
    return years
}

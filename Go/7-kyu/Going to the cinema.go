package kata

import "math"

func Movie(card, ticket int, perc float64) int {
    totalCostA := 0.0
    totalCostB := float64(card)
    currentTicketCost := float64(ticket)
    visits := 0

    for math.Ceil(totalCostB) >= totalCostA {
        visits++
        totalCostA += float64(ticket)
        currentTicketCost *= perc
        totalCostB += currentTicketCost
    }

    return visits
}

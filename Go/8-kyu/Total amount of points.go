package kata

import (
  "strconv"
  "strings"
)

func Points(games []string) int {
    points := 0

    for _, game := range games {
        score := strings.Split(game, ":")
        x, _ := strconv.Atoi(score[0])
        y, _ := strconv.Atoi(score[1])

        if x > y {
            points += 3 
        } else if x < y {
            points += 0 
        } else {
            points += 1 
        }
    }

    return points
}

package kata

import "math"

func WallPaper(l, w, h float64) string {
    numbers := []string{
        "zero", "one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine", "ten",
        "eleven", "twelve", "thirteen", "fourteen", "fifteen",
        "sixteen", "seventeen", "eighteen", "nineteen", "twenty",
    }

    if l == 0 || w == 0 || h == 0 {
        return numbers[0]
    }

    area := 2 * h * (l + w)

    area *= 1.15

    rollArea := 0.52 * 10

    rolls := int(math.Ceil(area / rollArea))

    if rolls > 20 {
        rolls = 20 
    }

    return numbers[rolls]
}

package kata

import (
    "strconv"
    "strings"
)

func parseCityData(town string, strng string) ([]float64, bool) {
    lines := strings.Split(strng, "\n")
    for _, line := range lines {
        parts := strings.Split(line, ":")
        if len(parts) != 2 || parts[0] != town {
            continue
        }
        monthlyData := strings.Split(parts[1], ",")
        values := make([]float64, 0, len(monthlyData))
        for _, item := range monthlyData {
            parts := strings.Split(strings.TrimSpace(item), " ")
            if len(parts) != 2 {
                continue
            }
            val, err := strconv.ParseFloat(parts[1], 64)
            if err != nil {
                return nil, false
            }
            values = append(values, val)
        }
        return values, true
    }
    return nil, false
}

func Mean(town string, strng string) float64 {
    values, found := parseCityData(town, strng)
    if !found {
        return -1.0
    }
    sum := 0.0
    for _, v := range values {
        sum += v
    }
    return sum / float64(len(values))
}

func Variance(town string, strng string) float64 {
    values, found := parseCityData(town, strng)
    if !found {
        return -1.0
    }
    mean := Mean(town, strng)
    var sum float64
    for _, v := range values {
        diff := v - mean
        sum += diff * diff
    }
    return sum / float64(len(values))
}

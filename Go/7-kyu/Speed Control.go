package kata

import "math"

func Gps(s int, x []float64) int {
    if len(x) <= 1 {
        return 0
    }
    
    maxSpeed := 0.0
    
    for i := 1; i < len(x); i++ {
        speed := (3600 * (x[i] - x[i-1])) / float64(s)
        if speed > maxSpeed {
            maxSpeed = speed
        }
    }
    
    return int(math.Floor(maxSpeed))
}

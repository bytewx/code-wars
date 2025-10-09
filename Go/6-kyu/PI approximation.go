package kata

import (
    "fmt"
    "math"
)

func IterPi(epsilon float64) (int, string) {
    var (
        pi4     float64 
        n       int     
        sign    = 1.0   
    )

    for {
        pi4 += sign / float64(2*n+1)
        approxPi := pi4 * 4
        if math.Abs(approxPi-math.Pi) < epsilon {
            return n + 1, fmt.Sprintf("%.10f", approxPi)
        }
        n++
        sign = -sign
    }
}

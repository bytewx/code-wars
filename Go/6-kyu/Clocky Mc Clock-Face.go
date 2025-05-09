package kata

import "fmt"

func WhatTimeIsIt(angle int) string {
    totalMinutes := angle * 2

    hours := (totalMinutes / 60) % 12
    minutes := totalMinutes % 60

    if hours == 0 {
        hours = 12
    }

    return fmt.Sprintf("%02d:%02d", hours, minutes)
}

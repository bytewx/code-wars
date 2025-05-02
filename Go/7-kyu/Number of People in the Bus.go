package kata

func Number(busStops [][2]int) int {
	total := 0
	for _, stop := range busStops {
		on := stop[0]
		off := stop[1]
		total += on - off
	}
	return total
}

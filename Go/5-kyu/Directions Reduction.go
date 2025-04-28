package kata

func DirReduc(arr []string) []string {
	opposite := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := []string{}

	for _, dir := range arr {
		n := len(stack)
		if n > 0 && opposite[dir] == stack[n-1] {
			stack = stack[:n-1] 
		} else {
			stack = append(stack, dir)
		}
	}

	return stack
}

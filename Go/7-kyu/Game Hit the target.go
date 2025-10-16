package kata

func Solution(mtrx [][]rune) bool {
	for _, row := range mtrx {
		var arrow, target = -1, -1
		for i, cell := range row {
			if cell == '>' {
				arrow = i
			} else if cell == 'x' {
				target = i
			}
		}
		if arrow != -1 && target != -1 && arrow < target {
			return true
		}
	}
	return false
}

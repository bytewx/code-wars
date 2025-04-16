package kata

func seatsInTheater(nCols int, nRows int, col int, row int) int {
    return (nRows - row) * (nCols - col + 1)
}

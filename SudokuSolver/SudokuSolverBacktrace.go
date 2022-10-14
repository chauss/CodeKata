package main

func SolveSudokuBacktrace(board *[9][9]uint8) bool {
	if !boardHasEmptyCell(*board) {
		return true
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[row][col] = uint8(candidate)
					if isBoardValid(*board) {
						if SolveSudokuBacktrace(board) {
							return true
						}
						board[row][col] = 0
					} else {
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

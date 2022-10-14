package main

func SolveSudoku(board [9][9]uint8) [9][9]uint8 {
	if !isBoardValid(board) {
		panic("Given board is not valid and can therefor not be solved!")
	}

	board_possibilities := [9][9][]uint8{}
	safety_counter := 0

	for boardHasEmptyCell(board) {

		// Step 1: Set all obvious solutions and create a possibility map
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[0]); col++ {
				if board[row][col] == 0 {
					board_possibilities = checkPossibilitesRecursively(&board, row, col, board_possibilities)
				}
			}
		}

		// Step 2: Analyse possibility map
		// Is there a field in (row, col, section) that is the only one with a missing number?

		printBoard(board)

		safety_counter++
		if safety_counter > 4 {
			break
		}
	}

	return board
}

func calcPossibleValues(board [9][9]uint8, row int, col int) []uint8 {
	possible_values := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Check all values in row
	for cur_col := 0; cur_col < len(board); cur_col++ {
		possible_values = removeValue(possible_values, board[row][cur_col])
	}

	// Check all values in col
	for cur_row := 0; cur_row < len(board); cur_row++ {
		possible_values = removeValue(possible_values, board[cur_row][col])
	}

	// Check all values in 3x3 section
	section_min_row := row / 3 * 3
	section_max_row := section_min_row + 2

	section_min_col := row / 3 * 3
	section_max_col := section_min_col + 2
	for cur_row := section_min_row; row <= section_max_row; row++ {
		for cur_col := section_min_col; row <= section_max_col; row++ {
			possible_values = removeValue(possible_values, board[cur_row][cur_col])
		}
	}

	return possible_values
}

func checkPossibilitesRecursively(board *[9][9]uint8, row int, col int, board_possibilities [9][9][]uint8) [9][9][]uint8 {
	possible_values := calcPossibleValues(*board, row, col)
	if len(possible_values) == 1 {
		board[row][col] = possible_values[0]
		board_possibilities = updateAffectedPositions(board, row, col, board_possibilities)
	} else {
		board_possibilities[row][col] = possible_values
	}

	return board_possibilities
}

func updateAffectedPositions(board *[9][9]uint8, row int, col int, board_possibilities [9][9][]uint8) [9][9][]uint8 {
	// Check affected row
	for cur_col := 0; cur_col < len(*board); cur_col++ {
		board_possibilities = checkPossibilitesRecursively(board, row, cur_col, board_possibilities)
	}

	// Check affected col
	for cur_row := 0; cur_row < len(board); cur_row++ {
		board_possibilities = checkPossibilitesRecursively(board, cur_row, col, board_possibilities)
	}

	// Check all values in 3x3 section
	section_min_row := row / 3 * 3
	section_max_row := section_min_row + 2

	section_min_col := row / 3 * 3
	section_max_col := section_min_col + 2
	for cur_row := section_min_row; row <= section_max_row; row++ {
		for cur_col := section_min_col; row <= section_max_col; row++ {
			board_possibilities = checkPossibilitesRecursively(board, cur_row, cur_col, board_possibilities)
		}
	}

	return board_possibilities
}

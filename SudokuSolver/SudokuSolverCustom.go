package main

import "fmt"

func SolveSudokuSophisticated(board [9][9]uint8) [9][9]uint8 {
	if !isBoardValid(board) {
		panic("Given board is not valid and can therefor not be solved!")
	}

	counter := 0

	last_iteration_board := [9][9]uint8{}

	for boardHasEmptyCell(board) {
		counter++
		board_possibilities := [9][9][]uint8{}

		// Step 1: Set all obvious solutions and create a possibility map
		board_has_been_updated := true
		for board_has_been_updated {
			board_possibilities = [9][9][]uint8{}

			fmt.Printf("Starting board analysis in iteration %d\n", counter)
			board_has_been_updated = false

			for row := 0; row < len(board); row++ {
				for col := 0; col < len(board[0]); col++ {
					if board[row][col] == 0 {
						possible_values := calcPossibleValues(board, row, col)
						if len(possible_values) == 1 {
							board[row][col] = possible_values[0]
							board_has_been_updated = true
							break
						} else {
							board_possibilities[row][col] = possible_values
						}
					}
				}
				if board_has_been_updated {
					break
				}
			}
		}

		if !shouldContinue(board, counter) {
			return board
		}

		// Step 2: Analyse possibility map
		// Is there a field in (row, col, section) that is the only one with a missing number?
		solutions := analysePossibilities(board, board_possibilities)
		fmt.Printf("Solutions: %d\n", solutions)
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[0]); col++ {
				if solutions[row][col] != 0 {
					board[row][col] = solutions[row][col]
					if !isBoardValid(board) {
						fmt.Printf("Wrong solution! (%d, %d) => %d\n", row, col, solutions[row][col])
					}
				}
			}
		}

		if !shouldContinue(board, counter) {
			return board
		}

		fmt.Printf("Board after %d iterations\n", counter)
		printBoard(board)

		if boardsAreEqual(board, last_iteration_board) {
			fmt.Println("Board didn't change in last iteration. Aborting!")
			break
		}

		for row := range board {
			last_iteration_board[row] = board[row]
		}
	}

	return board
}

func shouldContinue(board [9][9]uint8, counter int) bool {
	if !boardHasEmptyCell(board) && isBoardValid(board) {
		fmt.Printf("Board solved in %d iterations!\n", counter)
		printBoard(board)
		return false
	}
	if !isBoardValid(board) {
		fmt.Printf("Board moved to an invalid state. Aborting!")
		printBoard(board)
		return false
	}
	return true
}

func analysePossibilities(board [9][9]uint8, possibilities [9][9][]uint8) [9][9]uint8 {
	fmt.Println("Analysing possibilities!")
	solutions := [9][9]uint8{}

	// Analyse rows
	for row := 0; row < len(possibilities); row++ {
		cells_by_possible_solution := make(map[uint8][][]int)

		for col := 0; col < len(possibilities[0]); col++ {
			if board[row][col] != 0 {
				continue
			}
			for _, possible_solution := range possibilities[row][col] {
				if cells_by_possible_solution[possible_solution] == nil {
					cells_by_possible_solution[possible_solution] = [][]int{}
				}
				cells_by_possible_solution[possible_solution] = append(cells_by_possible_solution[possible_solution], []int{row, col})
			}
		}

		for solution, cells := range cells_by_possible_solution {
			if len(cells) == 1 {
				row := cells[0][0]
				col := cells[0][1]
				solutions[row][col] = solution
				fmt.Printf("Row - One solution is: (%d, %d) => %d\n", row, col, solution)
			}
		}
	}

	// Analyse columns
	for col := 0; col < len(possibilities[0]); col++ {
		cells_by_possible_solution := make(map[uint8][][]int)

		for row := 0; row < len(possibilities); row++ {
			if board[row][col] != 0 {
				continue
			}
			for _, possible_solution := range possibilities[row][col] {
				if cells_by_possible_solution[possible_solution] == nil {
					cells_by_possible_solution[possible_solution] = [][]int{}
				}
				cells_by_possible_solution[possible_solution] = append(cells_by_possible_solution[possible_solution], []int{row, col})
			}

		}

		for solution, cells := range cells_by_possible_solution {
			if len(cells) == 1 {
				row := cells[0][0]
				col := cells[0][1]
				solutions[row][col] = solution
				fmt.Printf("Col - One solution is: (%d, %d) => %d\n", row, col, solution)
			}
		}
	}

	// Check all 3x3 sections for duplicates
	for i := 0; i < len(possibilities); i += 3 {
		for j := 0; j < len(possibilities[0]); j += 3 {
			cells_by_possible_solution := make(map[uint8][][]int)

			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					if board[row][col] != 0 {
						continue
					}
					for _, possible_solution := range possibilities[row][col] {
						if cells_by_possible_solution[possible_solution] == nil {
							cells_by_possible_solution[possible_solution] = [][]int{}
						}
						cells_by_possible_solution[possible_solution] = append(cells_by_possible_solution[possible_solution], []int{row, col})
					}
				}
			}

			for solution, cells := range cells_by_possible_solution {
				if len(cells) == 1 {
					row := cells[0][0]
					col := cells[0][1]
					solutions[row][col] = solution
					fmt.Printf("3x3 - One solution is: (%d, %d) => %d\n", row, col, solution)
				}
			}
		}
	}

	return solutions
}

func calcPossibleValues(board [9][9]uint8, row int, col int) []uint8 {
	if board[row][col] != 0 {
		return []uint8{}
	}

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

	section_min_col := col / 3 * 3
	section_max_col := section_min_col + 2
	for cur_row := section_min_row; cur_row <= section_max_row; cur_row++ {
		for cur_col := section_min_col; cur_col <= section_max_col; cur_col++ {
			possible_values = removeValue(possible_values, board[cur_row][cur_col])
		}
	}

	// fmt.Printf("Possible Values for row=%d, col=%d => %d\n", row, col, possible_values)

	return possible_values
}

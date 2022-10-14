package main

import "fmt"

func printBoard(board [9][9]uint8) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < len(board); row++ {
		fmt.Print("| ")
		for col := 0; col < len(board[0]); col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
		}
		fmt.Println("|")
		if row == 2 || row == 5 {
			fmt.Println("+-------+-------+-------+")
		}
	}
	fmt.Println("+-------+-------+-------+")
}

func isBoardValid(board [9][9]uint8) bool {
	// Check all rows for duplicates
	for row := 0; row < len(board); row++ {
		counter := [10]uint8{}
		for col := 0; col < len(board[0]); col++ {
			counter[board[row][col]]++
		}
		if counterHasDuplicates(counter) {
			return false
		}
	}

	// Check all columns for duplicates
	for row := 0; row < len(board); row++ {
		counter := [10]uint8{}
		for col := 0; col < len(board[0]); col++ {
			counter[board[col][row]]++
		}
		if counterHasDuplicates(counter) {
			return false
		}
	}

	// Check all 3x3 sections for duplicates
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[0]); j += 3 {
			counter := [10]uint8{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if counterHasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func counterHasDuplicates(counter [10]uint8) bool {
	// Note: We start at idx := 1 because 0 means "not filled" and ignored
	for idx := 1; idx < len(counter); idx++ {
		if counter[idx] > 1 {
			return true
		}
	}
	return false
}

func boardHasEmptyCell(board [9][9]uint8) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == 0 {
				return true
			}
		}
	}
	return false
}

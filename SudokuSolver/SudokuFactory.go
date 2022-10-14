package main

func MediumSudokuBoard() [9][9]uint8 {
	return [9][9]uint8{
		{0, 0, 9, 3, 8, 0, 0, 0, 7},
		{0, 0, 0, 7, 6, 0, 0, 0, 0},
		{8, 2, 7, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 7, 8, 3, 2},
		{0, 0, 0, 0, 0, 4, 1, 9, 0},
		{2, 0, 8, 9, 3, 6, 7, 4, 5},
		{0, 0, 2, 1, 4, 5, 0, 0, 0},
		{1, 0, 0, 0, 0, 3, 0, 2, 9},
		{3, 7, 0, 2, 9, 0, 5, 0, 0},
	}
}

func EvilSudokuBoard() [9][9]uint8 {
	return [9][9]uint8{
		{3, 8, 0, 1, 0, 0, 0, 7, 0},
		{0, 0, 0, 0, 0, 4, 2, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 0, 0},
		{7, 5, 0, 3, 0, 0, 0, 8, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 3},
		{5, 6, 0, 0, 2, 0, 7, 0, 0},
		{0, 0, 9, 5, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 0},
	}
}
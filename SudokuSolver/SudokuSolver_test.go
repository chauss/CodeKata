package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveSudoku(t *testing.T) {
	board := [9][9]uint8{
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
	solved_board := SolveSudoku(board)
	printBoard(solved_board)
}

func TestCalPossibleValues(t *testing.T) {
	assertEqual := assert.New(t).Equal

	board := [9][9]uint8{
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

	possible_values := calcPossibleValues(board, 0, 0)
	fmt.Println(possible_values)
	assertEqual(possible_values, []uint8{4, 5, 6})
}

func TestBoardValidation(t *testing.T) {
	assertEqual := assert.New(t).Equal
	board := [9][9]uint8{
		{0, 0, 9, 3, 8, 0, 0, 0, 7}, // <- 9
		{0, 9, 0, 7, 6, 0, 0, 0, 0}, // <- 9
		{8, 2, 7, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 7, 8, 3, 2},
		{0, 0, 0, 0, 0, 4, 1, 9, 0},
		{2, 0, 8, 9, 3, 6, 7, 4, 5},
		{0, 0, 2, 1, 4, 5, 0, 0, 0},
		{1, 0, 0, 0, 0, 3, 0, 2, 9},
		{3, 7, 0, 2, 9, 0, 5, 0, 0},
	}
	isBoardValid := isBoardValid(board)
	assertEqual(false, isBoardValid)
}

func TestSliceUtils(t *testing.T) {
	assertEqual := assert.New(t).Equal
	slice := []uint8{}

	fmt.Println(slice)
	slice = addValue(slice, 5)
	slice = addValue(slice, 2)
	slice = addValue(slice, 5)
	slice = addValue(slice, 5)
	assertEqual(slice, []uint8{5, 2, 5, 5})
	slice = removeIndex(slice, 2)
	assertEqual(slice, []uint8{5, 2, 5})
	slice = removeValue(slice, 5)
	assertEqual(slice, []uint8{2})
}

func TestStuff(t *testing.T) {
	for i := 0; i < 9; i++ {
		lower := i / 3 * 3
		fmt.Printf("%d - %d\n", lower, lower+2)
	}
}

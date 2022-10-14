package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMediumSudokuCustom(t *testing.T) {
	assertTrue := assert.New(t).True

	solved_board := SolveSudokuSophisticated(MediumSudokuBoard())
	assertTrue(!boardHasEmptyCell(solved_board))
	assertTrue(isBoardValid(solved_board))
}

func TestSolveEvilSudokuCustom(t *testing.T) {
	assertTrue := assert.New(t).True

	solved_board := SolveSudokuSophisticated(EvilSudokuBoard())
	assertTrue(!boardHasEmptyCell(solved_board))
	assertTrue(isBoardValid(solved_board))
}

func TestSolveMediumSudokuBruteforce(t *testing.T) {
	assertTrue := assert.New(t).True

	board := MediumSudokuBoard()
	SolveSudokuBacktrace(&board)
	assertTrue(!boardHasEmptyCell(board))
	assertTrue(isBoardValid(board))
}

func TestSolveEvilSudokuBruteforce(t *testing.T) {
	assertTrue := assert.New(t).True

	board := EvilSudokuBoard()
	SolveSudokuBacktrace(&board)
	assertTrue(!boardHasEmptyCell(board))
	assertTrue(isBoardValid(board))
}

func TestCalPossibleValues(t *testing.T) {
	assertEqual := assert.New(t).Equal

	board := MediumSudokuBoard()

	possible_values := calcPossibleValues(board, 0, 0)
	assertEqual([]uint8{4, 5, 6}, possible_values)
	possible_values = calcPossibleValues(board, 1, 5)
	assertEqual([]uint8{1, 2, 9}, possible_values)
	possible_values = calcPossibleValues(board, 5, 1)
	assertEqual([]uint8{1}, possible_values)
	possible_values = calcPossibleValues(board, 8, 8)
	assertEqual([]uint8{1, 4, 6, 8}, possible_values)
	possible_values = calcPossibleValues(board, 4, 4)
	assertEqual([]uint8{2, 5}, possible_values)
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
	assertEqual([]uint8{5, 2, 5, 5}, slice)
	slice = removeIndex(slice, 2)
	assertEqual([]uint8{5, 2, 5}, slice)
	slice = removeValue(slice, 5)
	assertEqual([]uint8{2}, slice)

	slice2 := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 = removeValue(slice2, 9)
	assertEqual([]uint8{1, 2, 3, 4, 5, 6, 7, 8}, slice2)
}

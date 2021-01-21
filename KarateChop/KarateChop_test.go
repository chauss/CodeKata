package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfRecursive(t *testing.T) {
	karateChopTest(t, IndexOfRecursive)
}

func TestIndexOfLeftRight(t *testing.T) {
	karateChopTest(t, IndexOfLeftRight)
}

func karateChopTest(t *testing.T, indexOf func(int, []int) int) {
	assertEqual := assert.New(t).Equal

	t.Logf("\nEmpty Array")
	testSliceOne := []int{}
	assertEqual(-1, indexOf(1, testSliceOne))
	assertEqual(-1, indexOf(2, testSliceOne))
	assertEqual(-1, indexOf(-1, testSliceOne))

	t.Logf("\nArray with single item")
	testSliceTwo := []int{1}
	assertEqual(0, indexOf(1, testSliceTwo))
	assertEqual(-1, indexOf(0, testSliceTwo))
	assertEqual(-1, indexOf(2, testSliceTwo))

	t.Logf("\nArray with 4 items, all even")
	testSliceThree := []int{0, 2, 4, 6}
	assertEqual(0, indexOf(0, testSliceThree))
	assertEqual(1, indexOf(2, testSliceThree))
	assertEqual(2, indexOf(4, testSliceThree))
	assertEqual(3, indexOf(6, testSliceThree))
	assertEqual(-1, indexOf(-1, testSliceThree))
	assertEqual(-1, indexOf(1, testSliceThree))
	assertEqual(-1, indexOf(3, testSliceThree))
	assertEqual(-1, indexOf(5, testSliceThree))
	assertEqual(-1, indexOf(7, testSliceThree))

	t.Logf("\nArray with 6 items, including negative, all uneven")
	testSliceFour := []int{-1, 1, 3, 5, 7, 9}
	assertEqual(0, indexOf(-1, testSliceFour))
	assertEqual(1, indexOf(1, testSliceFour))
	assertEqual(2, indexOf(3, testSliceFour))
	assertEqual(3, indexOf(5, testSliceFour))
	assertEqual(4, indexOf(7, testSliceFour))
	assertEqual(5, indexOf(9, testSliceFour))
	assertEqual(-1, indexOf(-2, testSliceFour))
	assertEqual(-1, indexOf(0, testSliceFour))
	assertEqual(-1, indexOf(2, testSliceFour))
	assertEqual(-1, indexOf(4, testSliceFour))
	assertEqual(-1, indexOf(6, testSliceFour))
	assertEqual(-1, indexOf(8, testSliceFour))
	assertEqual(-1, indexOf(10, testSliceFour))
}

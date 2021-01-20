package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	assert := assert.New(t)

	t.Logf("\nEmpty Array")
	testSliceOne := []int{}
	assert.Equal(-1, IndexOf(1, testSliceOne))
	assert.Equal(-1, IndexOf(2, testSliceOne))
	assert.Equal(-1, IndexOf(-1, testSliceOne))

	t.Logf("\nArray with single item")
	testSliceTwo := []int{1}
	assert.Equal(0, IndexOf(1, testSliceTwo))
	assert.Equal(-1, IndexOf(0, testSliceTwo))
	assert.Equal(-1, IndexOf(2, testSliceTwo))

	t.Logf("\nArray with 4 items, all even")
	testSliceThree := []int{0, 2, 4, 6}
	assert.Equal(0, IndexOf(0, testSliceThree))
	assert.Equal(1, IndexOf(2, testSliceThree))
	assert.Equal(2, IndexOf(4, testSliceThree))
	assert.Equal(3, IndexOf(6, testSliceThree))
	assert.Equal(-1, IndexOf(-1, testSliceThree))
	assert.Equal(-1, IndexOf(1, testSliceThree))
	assert.Equal(-1, IndexOf(3, testSliceThree))
	assert.Equal(-1, IndexOf(5, testSliceThree))
	assert.Equal(-1, IndexOf(7, testSliceThree))

	t.Logf("\nArray with 6 items, including negative, all uneven")
	testSliceFour := []int{-1, 1, 3, 5, 7, 9}
	assert.Equal(0, IndexOf(-1, testSliceFour))
	assert.Equal(1, IndexOf(1, testSliceFour))
	assert.Equal(2, IndexOf(3, testSliceFour))
	assert.Equal(3, IndexOf(5, testSliceFour))
	assert.Equal(4, IndexOf(7, testSliceFour))
	assert.Equal(5, IndexOf(9, testSliceFour))
	assert.Equal(-1, IndexOf(-2, testSliceFour))
	assert.Equal(-1, IndexOf(0, testSliceFour))
	assert.Equal(-1, IndexOf(2, testSliceFour))
	assert.Equal(-1, IndexOf(4, testSliceFour))
	assert.Equal(-1, IndexOf(6, testSliceFour))
	assert.Equal(-1, IndexOf(8, testSliceFour))
	assert.Equal(-1, IndexOf(10, testSliceFour))
}

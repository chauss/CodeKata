package main

import "math"

// IndexOfLeftRight returns the index of number in list or -1 if list doesn't contain number
// The list must be sorted!
func IndexOfLeftRight(number int, list []int) int {
	left := 0
	right := len(list) - 1

	if len(list) == 0 {
		return -1
	}
	for left != right {
		binaryIndex := int(math.Ceil(float64(left+right) / 2))
		if list[binaryIndex] > number {
			right = binaryIndex - 1
		} else {
			left = binaryIndex
		}
	}
	if list[left] == number {
		return left
	}
	return -1
}

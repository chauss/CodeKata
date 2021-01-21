package karatechop

import (
	"math"
)

// IndexOfRecursive returns the index of number in list or -1 if list doesn't contain number
// The list must be sorted!
func IndexOfRecursive(number int, list []int) int {
	return indexOfRecursive(number, list, 0)
}

func indexOfRecursive(number int, list []int, initialIndex int) int {
	binaryIndex := int(math.Floor(float64(len(list)) / 2.0))

	if len(list) == 0 || list[0] > number || list[len(list)-1] < number {
		return -1
	} else if number < list[binaryIndex] {
		return indexOfRecursive(number, list[:binaryIndex], initialIndex)
	} else if number > list[binaryIndex] {
		return indexOfRecursive(number, list[binaryIndex:], initialIndex+binaryIndex)
	}
	return binaryIndex + initialIndex
}

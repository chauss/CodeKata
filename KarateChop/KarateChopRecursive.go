package recursive

import (
	"math"
)

// IndexOf returns the index of number in list or -1 if list doesn't contain number
// The list must be sorted!
func IndexOf(number int, list []int) int {
	return indexOf(number, list, 0)
}

func indexOf(number int, list []int, initialIndex int) int {
	binaryIndex := int(math.Floor(float64(len(list)) / 2.0))

	if len(list) == 0 || list[0] > number || list[len(list)-1] < number {
		return -1
	} else if number < list[binaryIndex] {
		return indexOf(number, list[:binaryIndex], initialIndex)
	} else if number > list[binaryIndex] {
		return indexOf(number, list[binaryIndex:], initialIndex+binaryIndex)
	}
	return binaryIndex + initialIndex
}

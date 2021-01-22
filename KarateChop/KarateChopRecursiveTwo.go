package main

// IndexOfRecursiveTwo returns the index of number in list or -1 if list doesn't contain number
// The list must be sorted!
func IndexOfRecursiveTwo(number int, list []int) int {
	if len(list) == 0 {
		return -1
	}

	left := 0
	right := len(list) - 1
	return indexOfRecursiveTwo(number, list, left, right)
}

func indexOfRecursiveTwo(number int, list []int, left int, right int) int {
	binaryIndex := (right + left) / 2

	if left > right {
		return -1
	} else if number < list[binaryIndex] {
		return indexOfRecursiveTwo(number, list, left, binaryIndex-1)
	} else if number > list[binaryIndex] {
		return indexOfRecursiveTwo(number, list, binaryIndex+1, right)
	}
	return binaryIndex
}

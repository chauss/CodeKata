package main

func addValue(slice []uint8, value uint8) []uint8 {
	return append(slice, value)
}

func removeValue(slice []uint8, value uint8) []uint8 {
	new_slice := []uint8{}

	for index := 0; index < len(slice); index++ {
		if slice[index] != value {
			new_slice = append(new_slice, slice[index])
		}
	}
	return new_slice
}

func removeIndex(slice []uint8, index int) []uint8 {
	// copy last element to index
	slice[index] = slice[len(slice)-1]
	// cut of last element
	slice = slice[:len(slice)-1]
	return slice
}

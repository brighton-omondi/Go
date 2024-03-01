package piscine

func ShoppingListSort(slice []string) []string {
	// Bubble sort implementation
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				// Swap elements
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

package piscine

func DescendAppendRange(max, min int) []int {
	// Check if max is inferior or equal to min
	if max <= min {
		return []int{} // Return an empty slice
	}

	// Initialize an empty slice to store the values
	var result []int

	// Iterate from max down to min+1 (excluding min)
	for i := max; i > min; i-- {
		// Append the current value to the slice
		result = append(result, i)
	}

	return result
}

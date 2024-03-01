package piscine

func CollatzCountdown(start int) int {
	// If start is 0 or negative, return -1
	if start <= 0 {
		return -1
	}

	// Initialize the count of steps
	steps := 0

	// Loop until start becomes 1
	for start != 1 {
		// If start is even, divide it by 2
		if start%2 == 0 {
			start /= 2
		} else {
			// If start is odd, multiply by 3 and add 1
			start = 3*start + 1
		}
		// Increment the step count
		steps++
	}

	// Return the number of steps
	return steps
}

package piscine

func Unmatch(a []int) int {
	var count int // Assuming the range of integers and adjust as needed

	// Iterate through the slice and count the frequency of each element
	for _, num := range a {
		count = 0
		for _, v := range a {
			if v == num {
				count++
			}
		}
		if count%2 != 0 {
			return num
		}
	}

	// If no unmatched element is found, return -1
	return -1
}

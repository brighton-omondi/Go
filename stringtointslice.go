package piscine

func StringToIntSlice(str string) []int {
	var nums []int

	// Loop through each character in the string
	for _, char := range str {
		nums = append(nums, int(char)) // Convert rune to int and append to slice
	}

	return nums
}

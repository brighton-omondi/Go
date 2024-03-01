package piscine

func Abort(a, b, c, d, e int) int {
	// Put the values into a slice
	arr := []int{a, b, c, d, e}
	Sort(arr)

	midean := len(arr) / 2
	// Return the middle value
	return arr[midean]
}

func Sort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

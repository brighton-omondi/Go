package piscine

import "fmt"

func PodiumPosition(podium [][]string) [][]string {
	// Correcting the positions based on the length of each place string
	for i := 0; i < len(podium); i++ {
		position := fmt.Sprintf("%d%s", i+1, suffix(i+1))
		podium[i][0] = position + " Place"
	}
	return podium
}

func suffix(num int) string {
	if num == 1 {
		return "st"
	} else if num == 2 {
		return "nd"
	} else if num == 3 {
		return "rd"
	}
	return "th"
}

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	programName := "main"
	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

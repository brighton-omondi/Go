package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	EvenMsg = "I have an even number of arguments\n"
	OddMsg  = "I have an odd number of arguments\n"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func lengthOfArg() int {
	return len(os.Args) - 1
}

func main() {
	if isEven(lengthOfArg()) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

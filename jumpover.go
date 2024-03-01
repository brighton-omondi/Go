package piscine

import (
	"github.com/01-edu/z01"
)

func JumpOver(str string) {
	if len(str) == 0 {
		z01.PrintRune('\n')
		return
	}

	for i := 2; i < len(str); i += 3 {
		z01.PrintRune(rune(str[i]))
	}

	if len(str) < 3 || len(str)%3 != 0 {
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}

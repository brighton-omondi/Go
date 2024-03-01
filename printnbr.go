package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	sign := 1
	if n < 0 {
		sign = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		number := (n / 10) * sign
		if number != 0 {
			PrintNbr(number)
		}
		k := (n%10)*sign + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}

package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for fourth := '9'; fourth >= '0'; fourth-- {
		for third := '9'; third >= '0'; third-- {
			for second := '9'; second >= '0'; second-- {
				for first := '9'; first >= '0'; first-- {
					if fourth > second || (fourth == second && third > first) {
						z01.PrintRune(fourth)
						z01.PrintRune(second)
						z01.PrintRune(' ')
						z01.PrintRune(third)
						z01.PrintRune(first)

						separator := ", "
						for _, x := range separator {
							if !(fourth == '0' && third <= '1' && second == '0' && first <= '0') {
								z01.PrintRune(x)
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	DescendComb()
	z01.PrintRune('\n')
}

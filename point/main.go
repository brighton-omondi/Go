package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, xter := range s {
		z01.PrintRune(xter)
	}
}

func printNum(n int) {
	first := n / 10
	second := n % 10

	digit1 := '0'
	for i := 0; i < first; i++ {
		digit1++
	}
	z01.PrintRune(digit1)

	digit2 := '0'
	for i := 0; i < second; i++ {
		digit2++
	}
	z01.PrintRune(digit2)
}

func main() {
	points := &point{}

	setPoint(points)
	printStr("x = ")
	printNum(points.x)
	printStr(", y = ")
	printNum(points.y)
	z01.PrintRune('\n')
}

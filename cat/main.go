package main

import (
	"io"
	"os"
)

func cat() {
	// Copy standard input to standard output
	io.Copy(os.Stdout, os.Stdin)
}

func main() {
	cat()
}

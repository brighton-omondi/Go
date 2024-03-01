package main

import (
	"fmt"
	"os"
)

func comcheck(args ...string) {
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

func main() {
	comcheck(os.Args[1:]...)
}

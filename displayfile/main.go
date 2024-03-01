package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arglen := len(os.Args)

	if arglen > 2 {
		fmt.Println("Too many arguments")
	} else if arglen == 1 {
		fmt.Println("File name missing")
	} else {
		data, _ := ioutil.ReadFile(os.Args[1])
		fmt.Printf("%s", string(data))
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //Uninitialized variables are initialized to their zero values
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

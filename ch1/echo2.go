package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""                  //Types will be implicitly determined
	for _, arg := range os.Args[1:] { //_ is a blank identifier, this for statement is diff from init;cond;post loop
		s += sep + arg
		sep = " "
	}

}

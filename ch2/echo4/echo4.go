//using flags to print command line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //.Bool returns pointer to the flag
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse() //To set default values
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

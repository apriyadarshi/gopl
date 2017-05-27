//dup3 reads the whole file  and then splits it by lines to count the duplicates.
//Only those lines which are repeated 2 or more times are shown.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
		}
		for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

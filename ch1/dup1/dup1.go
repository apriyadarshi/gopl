package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //make makes a new empty map string -> int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //An infinite loop
		counts[input.Text()]++
		// Errors from input. Err ignored ??
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

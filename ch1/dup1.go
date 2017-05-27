//Dup1 prints the text of each line that apeears more than once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      //make makes a new empty map string -> int
	input := bufio.NewScanner(os.Stdin) //NewScanner is a type in bufio package
	for input.Scan() {                  //An infinite loop
		counts[input.Text()]++
		// Errors from input. Err ignored ??
		for line, n := range counts { //Remember: Order of map iteration is random i.e. not in the same order as entered
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

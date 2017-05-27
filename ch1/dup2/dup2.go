//Dup2 prints the count and text of lines that appear more than once in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileNameWithCounts := make(map[string]map[string]int) //map: string -> (string -> int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts) //Since a compiled lang, a func maybe called before declaring it
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg) //f=open file; err=nil if file opened successfully
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			fileNameWithCounts[arg] = counts
			f.Close()
		}
	}
	for fileName, mapLineCount := range fileNameWithCounts {
		for line, n := range mapLineCount { // Notice that range works differently for maps and arrays: for arrays, it gives (index, value) but for maps it gives (key, value)
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) { //Copy of reference is passed in second argument.
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		//Ignoring potential errors from input.Err()
	}
}

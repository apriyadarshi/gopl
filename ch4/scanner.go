package ch4

import (
	"bufio"
	"log"
	"os"
)

//Punctutation marks are not handled. So "Hello.", "Hello!" and Hello are different words
func WordFreq(file string) map[string]int {

	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	result := make(map[string]int)
	for scanner.Scan() {
		result[scanner.Text()]++
	}

	return result
}

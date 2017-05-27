package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	var n = flag.Int("c", 0, "checksum to use 256, 384 or 512")

	flag.Parse()

	input := os.Args[1]
	//check if flag is valid
	switch *n {
	case 256:
		fmt.Printf("sha256 Sum of %s : %x", input, sha256.Sum256([]byte(input)))
	case 384:
		fmt.Printf("sha384 Sum of %s : %x", input, sha512.Sum384([]byte(input)))
	case 512:
		fmt.Printf("sha512 Sum of %s : %x", input, sha512.Sum512([]byte(input)))
	default:
		*fmt.Printf("Invalid flag %d. Please check valid values ushing -help", *n)
	}
}

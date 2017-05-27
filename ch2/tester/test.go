package main

import (
	"fmt"
	"gopl/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Printf("Setbits for the uint %d = %d", x, popcount.PopCount3(x))
}

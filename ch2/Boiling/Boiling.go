//Boiling prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0 //package level constant

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c) //%g prints the min no of digits to identify a value uniquely
}

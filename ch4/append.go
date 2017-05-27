//Appending elements at the end of slices
package ch4

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // Compare length with the length of underlying array
		z = x[:zlen] //Notice that RHS is a valid statement even when zlen > len(x)
	} else {
		//Allocate new array. Double  the size for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //Copies x in z
	}
	z[len(x)] = y
	return z
}

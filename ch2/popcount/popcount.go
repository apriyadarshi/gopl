//For an unsigned int's binary representation, popcount counts the number of bits that are 1 (also called set bits).
package popcount

var pc [256]byte //pc[i] is the popcount of i. A slice of 256 length each 1 byte long

//Init function can't be called but automatically executes for this file, when the program starts
//Purpose : Initializing a complex variable. In this case pc
func init() {
	for i := range pc { ///Same as for index, _
		pc[i] = pc[i/2] + byte(i&1) //Logic - Rightshift bitwise (div by 2) and add 1 bit if 1st bit from right is 1
	}
}

//Bug - if //Same as for index,_ := range pc typed, coloring breaks

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + //logic: Right shift ; byte(<a64bit number>) gets first 8 bytes from right
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

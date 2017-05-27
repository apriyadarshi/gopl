package ch4

import "gopl/ch2/popcount"

//This function counts the number of bits different in two SHA256 hashes
func DiffBits(sha1, sha2 [32]uint8) int {
	//Logic - Take bitwise XOR. Count no of 1's in result using popcount

	var xor [32]uint8
	var count int
	for i := range sha1 {
		xor[i] = sha1[i] ^ sha2[i]
		count += popcount.PopCountInt8(xor[i])
	}
	return count
}

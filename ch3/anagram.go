//Functions to check whether two strings are anagrams of each other
package ch3

import "bytes"

//Inefficient implementation w/o using maps
func Anagram1(s1, s2 string) bool { //both parameters are string
	b1, b2 := []byte(s1), []byte(s2)

	if len(b1) != len(b2) {
		return false
	} else {
		for i := range b1 {
			if bytes.Count(b1, b1[i:i+1]) != bytes.Count(b2, b1[i:i+1]) { //Double counting
				return false
			}
		}
	}
	return true
}

//Efficient one can be done using maps

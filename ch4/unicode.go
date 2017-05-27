package ch4

import (
	"unicode"
	"unicode/utf8"
)

//Inplace function that collapses adjacent spaces in a single ASCII space
func collapseSpaces(b []byte) []byte {
	var i, j int
	for i, j = 0, 0; j < len(b); {
		r1, w1 := utf8.DecodeRune(b[j:])
		if j+w1 < len(b) {
			r2, _ := utf8.DecodeRune(b[j+w1:])
			if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
				i = i //Don't increase i. Space will be printed when r1=space, r2 =nonspace
				j = j + w1
			} else {
				if unicode.IsSpace(r1) {
					utf8.EncodeRune(b[i:], (' '))
				} else {
					utf8.EncodeRune(b[i:], r1)
				}
				j = j + w1
				i = i + w1
			}
		} else {
			if unicode.IsSpace(r1) {
				utf8.EncodeRune(b[i:], (' '))
			} else {
				utf8.EncodeRune(b[i:], r1)
			}
			j = j + w1
			i = i + w1
		}
	}
	return b[0:i]
}

package ch4

import "unicode/utf8"

//Inplace reverse using slice
func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*Inplace using array pointer
Issue : type of *[10]int is different fom type of *[11]int
So, this can only be done for a fixed array size*/
//Both a and *a work
//Note: If p = &[1 2 3], p[1] will successfully give 2
func rev2(a *[5]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

//Reversing a byte slice that is UTF-8 encoded in-place.
//Reverse the characters i.e. reverse the runes
func revRunes(b []byte) {
	//Pass 1: Loop through runes. If a rune is multi-byte, reverse it
	for i := 0; i < len(b); {
		_, w := utf8.DecodeRune(b[i:])
		if w > 1 {
			revBytes(b[i : i+w])
		}
		i += w
	}

	//Pass 2: reverse the whole slice
	revBytes(b)
}

func revBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

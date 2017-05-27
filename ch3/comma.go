package ch3

import (
	"bytes"
	"strings"
)

//comma inserts comma in a non-negative decimal integer
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//comma 2 inserts comma in a non-negative decimal number
func Comma2(s string) string {

	dotI := strings.LastIndex(s, ".")

	if dotI != -1 {
		return comma(s[:dotI]) + s[dotI:]
	} else {
		return comma(s[:dotI])
	}
}

//Non-recursive bytes.Buffer version. ASSUMPTION :ASCII numbers
//Negative numbers allowed. Floating point numbers allowed
func Comma3(s string) string {
	b := []byte(s)

	//Handle sign
	neg := b[0] == byte('-')

	//Handle floats
	dotI := bytes.LastIndexAny(b, ".") // one left to dot

	var lDecI, fDecI int
	if dotI < 0 {
		lDecI = len(b) - 1 //len returns number of bytes. Fine here as ASCII characters take 1 bit only
	} else {
		lDecI = dotI - 1
	}

	if neg {
		fDecI = 1
	}

	decLen := lDecI - fDecI + 1
	if decLen < 4 {
		return s
	}

	//create output : neg + chunks in reverse + dot + fractional part
	var buf bytes.Buffer
	if neg {
		buf.WriteByte('-')
	}
	for i := fDecI; i < lDecI; {
		if i == fDecI {
			inc := 3
			if decLen%3 != 0 {
				inc = decLen % 3
			}
			buf.Write(b[fDecI : fDecI+inc])
			i = i + inc
		} else {
			buf.WriteByte(',')
			buf.Write(b[i : i+3])
			i = i + 3
		}
	}

	if dotI >= 0 {
		buf.Write(b[dotI:])
	}

	return buf.String()

}

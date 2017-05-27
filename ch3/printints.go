//..intsToString is like fmt.Sprint(values) but adds commas
package ch3

import (
	"bytes"
	"fmt"
)

func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v) //Fprintf formats as per specifier and writes to the first argument
	}
	buf.WriteByte(']')
	return buf.String()
}

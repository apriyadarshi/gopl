package expand

import (
	"strings"
)

//Replaces $foo with f("foo")
func Expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

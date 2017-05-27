package ch5

import (
	//"fmt"
	"strings"
)

func join(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

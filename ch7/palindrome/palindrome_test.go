package palindrome

import (
	//"fmt"
	//"sort"
	"testing"
)

type intset []int

func (ints intset) Len() int {
	return len(ints)
}

func (ints intset) Less(i, j int) bool {
	return ints[i] < ints[j]
}

func (ints intset) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func TestPalindrome(t *testing.T) {
	set := intset([]int{1, 2, 3, 4, 3, 2, 1})
	if !IsPalindrome(set) {
		t.Error("expected: palindrome; got: not palindrome")
	}
}

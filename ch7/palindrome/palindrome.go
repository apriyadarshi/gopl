package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	if s.Len() < 1 {
		return true
	}
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

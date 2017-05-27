package ch4

import (
	"unicode"
)

//charcount to count digits, letters etc
func charcount(s string) []map[rune]int {

	letters, digits, puncts, spaces := make(map[rune]int), make(map[rune]int), make(map[rune]int), make(map[rune]int)

	for _, r := range s {
		//runes[r]++ //Also includes the invalidchar count
		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsDigit(r) {
			digits[r]++
		}
		if unicode.IsSpace(r) {
			spaces[r]++
		}
		if unicode.IsPunct(r) {
			puncts[r]++
		}
	}

	return []map[rune]int{letters, digits, puncts, spaces}
}

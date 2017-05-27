package bitarray

import (
	"bytes"
	"fmt"
	"gopl/ch2/popcount"
)

const BITS = 32 << (^(uint(0)) >> 63) //Testing machine bits

//A IntSet is a set of small non-negative empty integers.
//Its zero value returns the empty set.
type IntSet struct {
	words []uint //Hidden.Can't be accessed by clients outside this package
}

//Has checks whether an integer is in the set or not
func (s *IntSet) Has(x int) bool {
	word, bit := x/BITS, uint(x%BITS)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add adds the non-negative value of x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/BITS, uint(x%BITS)
	for word >= len(s.words) {
		//Increase size of s.words adding zero values to end to
		//add word at the correct position.
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/BITS, uint(x%BITS)
	if word < len(s.words) {
		s.words[word] &= ^(uint(1) << bit)
		//^ performs biwise negation for uint. uint(1) is written as negation works diff for int(2's comp) and uint
	}
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

//UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//IntersectWith sets s to the union of s and t
func (s *IntSet) Intersect(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i >= len(t.words) { //for words not in t, set s==0
			s.words[i] &= 0
		} else {
			s.words[i] &= t.words[i]
		}
	}
}

//DifferenceWith finds s - t
//For each bit diff. a-b = ab` => a and not b
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &^= (t.words[i]) //same as and with negation of second
		}
	}
}

//Symmetric DifferenceWith finds elemnts which are in s or in t but not in both
//For each bit diff. a XOR b i.e. a^b
func (s *IntSet) SymDiffWith(t *IntSet) {
	for i := 0; i < len(t.words); i++ {
		if i < len(s.words) {
			s.words[i] ^= (t.words[i])
		} else {
			s.words = append(s.words, t.words[i])
		}
	}

}

//String returns the set as a string of the form "{1 2 3}"
//Helps printing using fmt.Printf("%s", ...)
//Imp: IntSet value doesn't have this method. IntSet pointer has
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITS; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*BITS+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//Number of elements in set
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		count += popcount.PopCountUInt(word)
	}
	return count
}

//Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	dstWords := make([]uint, len(s.words), cap(s.words))
	copy(dstWords, s.words) //Imp: Copy for slices copies in destination min(len(dst),
	c := IntSet{words: dstWords}
	return &c
}

//Return a copy of the actual slice
func (s *IntSet) Elems() []uint {
	dstWords := make([]uint, len(s.words), cap(s.words))
	copy(dstWords, s.words)
	return dstWords
}

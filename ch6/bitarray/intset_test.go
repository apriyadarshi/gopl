package bitarray

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	set := IntSet{[]uint{}}

	if i := set.Len(); i != 0 {
		t.Errorf("Len : expected : %d, got: %d", 0, i)
	}

	set.Add(3)
	set.Add(4)
	set.Add(5)

	if i := set.Len(); i != 3 {
		t.Errorf("Len : expected : %d, got: %d", 3, i)
	}

}

func TestRemove(t *testing.T) {
	set := IntSet{[]uint{}}

	set.Remove(1) //No error here => Removing from empty error doesn't throw panic which is an expected behavioyr

	set.Add(3)
	set.Add(4)
	set.Add(5)

	set.Remove(5)
	fmt.Printf("%s", &set)

	if set.Has(5) {
		t.Error("Element not removed")
	}
}

func TestClear(t *testing.T) {
	set := IntSet{[]uint{}}
	set.Add(3)
	set.Add(4)
	set.Add(5)
	fmt.Printf("Before clearing: %s\n", &set)
	set.Clear()
	fmt.Printf("After clearing: %s\n", &set)

	if set.Len() != 0 {
		t.Error("Elements still present: %s", &set)
	}
}

func TestCopy(t *testing.T) {
	set := IntSet{[]uint{}}
	set.Add(3)
	set.Add(4)
	set.Add(5)
	cset := set.Copy()
	fmt.Printf("Set before clearing set: %s\n", &set)
	fmt.Printf("CopySet before clearing set: %s\n", cset)
	set.Clear()
	fmt.Printf("Set after clearing set: %s\n", &set)
	fmt.Printf("CopySet after clearing set: %s\n", cset)
	if cset.Len() != 3 {
		t.Error("Copying doesn't create true copy. Clearing src set clears dst set")
	}
}

func TestAddAll(t *testing.T) {
	set := IntSet{[]uint{}}
	set.AddAll(3, 4, 5)
	if !(set.Len() == 3 && set.Has(3) && set.Has(4) && set.Has(5)) {
		t.Errorf("addall. expected: %s, got:%s", "{3, 4, 5}", set)
	}
}

func TestIntersect(t *testing.T) {
	setA, setB := IntSet{[]uint{}}, IntSet{[]uint{}}

	//Case: len(setA) > len(setB) & no
	setA.AddAll(1, 2, 64, 65, 128, 129)
	setB.AddAll(1, 65)

	setA.Intersect(&setB)
	fmt.Printf("A was: %s", &setA)
	fmt.Printf("%s", &setB)
	if !(setA.Len() == 2 && setA.Has(1) && setA.Has(65)) {
		t.Errorf("intersect: expected %s, got: %s", "{1, 65}", &setA)
	}

	//Case: len(setA) < len(setB)
	setB.AddAll(2, 128, 129, 257, 657)
	setA.Intersect(&setB)
	if !(setA.Len() == 2 && setA.Has(1) && setA.Has(65)) {
		t.Errorf("intersect: expected %s, got: %s", "{1, 65}", &setA)
	}

}

func TestDifferenceWith(t *testing.T) {
	setA, setB := IntSet{[]uint{}}, IntSet{[]uint{}}

	//Case: len(setA) > len(setB) & no
	setA.AddAll(1, 2, 64, 65, 128, 129)
	setB.AddAll(1, 65, 66)

	setA.DifferenceWith(&setB)
	fmt.Printf("A was: %s\n", &setA)
	fmt.Printf("B was:%s\n", &setB)
	if !(setA.Len() == 4 && setA.Has(2) && setA.Has(64) && setA.Has(128) && setA.Has(129)) {
		t.Errorf("difference: expected %s, got: %s", "{2, 64, 128, 129}", &setA)
	}

	setB.AddAll(356, 357, 128)
	setA.DifferenceWith(&setB)
	fmt.Printf("A was: %s\n", &setA)
	fmt.Printf("B was:%s\n", &setB)
	if !(setA.Len() == 3 && setA.Has(2) && setA.Has(64) && setA.Has(129)) {
		t.Errorf("difference: expected %s, got: %s", "{1, 64, 128}", &setA)
	}
}

func TestSymDifferenceWith(t *testing.T) {
	setA, setB := IntSet{[]uint{}}, IntSet{[]uint{}}

	//Case: len(setA) > len(setB) & no
	setA.AddAll(1, 2, 64, 65, 128, 129)
	setB.AddAll(1, 65, 66)

	fmt.Printf("A was: %s\n", &setA)
	fmt.Printf("B was:%s\n", &setB)
	setA.SymDiffWith(&setB)
	fmt.Printf("A is: %s\n", &setA)
	//fmt.Printf("B was:%s\n", &setB)

	if !(setA.Len() == 5 && setA.Has(2) && setA.Has(64) && setA.Has(66) && setA.Has(128) && setA.Has(129)) {
		t.Errorf("symmetric difference: expected %s, got: %s", "{2,64,66,128,129}", &setA)
	}

	setB.AddAll(356, 357, 128)
	fmt.Printf("A was: %s\n", &setA) //A was: {2 64 66 128 129}
	fmt.Printf("B was:%s\n", &setB)  //B was:{1 65 66 128 356 357}
	setA.SymDiffWith(&setB)
	fmt.Printf("A is: %s\n", &setA)
	if !(setA.Len() == 7 && setA.String() == "{1 2 64 65 129 356 357}") {
		t.Errorf("symmetric difference: expected %s, got: %s", "{1 2 64 65 129 356 357}", &setA)
	}
}

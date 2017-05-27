package ch5

import "testing"

func TestMin(t *testing.T) {
	nums := []int{5, 4, 2, 1, 7, 9}
	if min, _ := min(nums...); min != 1 {
		t.Errorf("incorrect min. expected : %d found: %d", 1, min)
	}
}

func TestMax(t *testing.T) {
	nums := []int{5, 4, 2, 1, 7, 9}
	if max, _ := max(nums...); max != 9 {
		t.Errorf("incorrect max. expected : %d found: %d", 9, min)
	}
}

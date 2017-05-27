package ch5

import "testing"

func TestJoin(t *testing.T) {
	if r, e := join(" ", "Hello!", "How", "are", "you?"), "Hello! How are you?"; r != e {
		t.Errorf("ch5.join : expected : '%s' found '%s'", e, r)
	}
}

package expand

import "testing"

func TestExpand(t *testing.T) {
	s := Expand("Who was here? $foo.\nWill someone tell? $foo ", Dummy)
	corr := "Who was here? !Dummy was here!.\nWill someone tell? !Dummy was here! "
	if s != corr {
		t.Errorf("expand failed: \nExpected: %s\nGot: %s\n", corr, s)
	}
}

func Dummy(s string) string {
	return "!Dummy was here!"
}

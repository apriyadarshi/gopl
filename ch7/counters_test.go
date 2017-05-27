package ch5

import (
	"fmt"
	"os"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	fmt.Fprint(&c, "Hello! How are we today")
	if int(c) != 5 {
		t.Errorf("wordcounter: expected: 5 got: ", int(c))
	}
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	fmt.Fprint(&c, "A: Hello! How are we today\nB: Fine! How are you?")
	if int(c) != 2 {
		t.Errorf("linecounter: expected: 2 got: ", int(c))
	}
}

func TestCountingWriter(t *testing.T) {
	cw, n := CountingWriter(os.Stderr)
	n1, _ := cw.Write([]byte("Hi"))
	n2, _ := cw.Write([]byte("Dude"))
	if *n != int64(n1+n2) {
		t.Errorf("Incorrect number of bytes counted. Expected count: %d, Actual Count: %d", n1+n2, *n)
	}
}

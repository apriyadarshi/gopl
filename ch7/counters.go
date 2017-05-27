package ch5

import (
	"bufio"
	"bytes"
	"io"
)

//Type bytecounter counts the number of bytes
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) //Int converted to bytecounter
	return len(p), nil
}

//type wordcounter counts the number of words in input
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var count int

	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	*c = WordCounter(count)
	return len(p), nil
}

//type LineCounter int
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var count int

	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		count++
	}

	*c = LineCounter(count)
	return len(p), nil
}

//
type CWriter struct {
	writer io.Writer
	count  int64
}

func (cw *CWriter) Write(p []byte) (int, error) {
	i, err := cw.writer.Write(p)
	cw.count += int64(i)
	return i, err
}

//Returns a new writer wrapping the original.
//Second return value always gives the number of bytes writtern.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var n int64
	cw := CWriter{writer: w, count: n}
	return &cw, &cw.count
}

package ch5

import (
	"bytes"
	"io"
)

type Reader struct {
	r io.Reader
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func NewReader(s string) *Reader {
	r := Reader{r: bytes.NewBuffer([]byte(s))}
	return &r
}

/*A LimitedReader reads from R but limits the amount of data returned to just N bytes.
Each call to Read updates N to reflect the new amount remaining.
Read returns EOF when N <= 0 or when the underlying R returns EOF.
*/

type LimitedReader struct {
	r io.Reader
	n int64
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	i, err := lr.r.Read(p)
	lr.n -= int64(i)
	if err != nil {
		return i, err
	}
	if lr.n <= 0 {
		return i, io.EOF
	}

	return i, nil
}

func LimitReader(reader io.Reader, num int64) io.Reader {
	lr := LimitedReader{r: reader, n: num}
	return &lr
}

package filter

import (
	"bytes"
	"io"
	"io/ioutil"
)

// NewReadCloser ...
func NewReadCloser(rc io.ReadCloser, f RuneMap) io.ReadCloser {
	return &Reader{rc, rc, f}
}

// NewReader ...
func NewReader(r io.Reader, f RuneMap) io.Reader {
	return &Reader{r, ioutil.NopCloser(r), f}
}

// Reader uses bytes.Map to transform the input
// from an io.Reader to alter an inbound stream
type Reader struct {
	r       io.Reader
	c       io.Closer
	mapping RuneMap
}

// Read checks for changes to the slice, then rewrites the slice
// if any are found
func (f *Reader) Read(p []byte) (int, error) {
	n, err := f.r.Read(p)
	// nothing is changed by mapping
	if f.isSame(p[:n]) {
		return n, err
	}
	// since there are differences, map them out
	data := bytes.Map(f.mapping, p[:n])
	// copy the data over
	for i := range data {
		p[i] = data[i]
	}
	return len(data), nil
}

// Close meets the io.Closer interface
func (f *Reader) Close() error {
	return f.c.Close()
}

// isSame verifies that the slice is unchanged (optimization for Read)
func (f *Reader) isSame(p []byte) bool {
	// see if any runes were dropped
	for _, r := range bytes.Runes(p) {
		if f.mapping(r) != r {
			return false
		}
	}
	return true
}

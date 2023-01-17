package extraio

import (
	"io"
)

// DiscardReader is an io.Reader that discard all read bytes from the underlying
// reader.
//
// Note that its Read method returns zero byte count. Some io.Reader client
// implementations return io.ErrNoProgress error when many calls to Read have
// failed to return any data or error.
type DiscardReader struct {
	reader io.Reader
}

// NewDiscardReader returns a new reader that discard all reads from r.
func NewDiscardReader(r io.Reader) *DiscardReader {
	return &DiscardReader{r}
}

// Read implements the io.Reader interface. It reads from the underlying
// io.Reader but always returns zero byte count.
func (d *DiscardReader) Read(p []byte) (int, error) {
	_, err := d.reader.Read(p)
	return 0, err
}

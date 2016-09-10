package gzip

import (
	"bytes"
	"compress/gzip"
)

// GetGZIPFileReader returns a reader to an in-memory
// gzip archive containing `content`
func GetGZIPFileReader(content []byte) (*bytes.Reader, error) {
	var buf bytes.Buffer

	w := gzip.NewWriter(&buf)

	if _, err := w.Write(content); err != nil {
		return nil, err
	}

	w.Close()

	return bytes.NewReader(buf.Bytes()), nil
}

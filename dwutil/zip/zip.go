package zip

import (
	"archive/zip"
	"bytes"
)

// FileContent represents the name and content of a file
type FileContent struct {
	Name string
	Body []byte
}

// FileCollection is a collection of files
type FileCollection []FileContent

// GetZIPFileReader returns a reader to an in-memory
// zip archive containing `content`
func GetZIPFileReader(content FileCollection) (*bytes.Reader, error) {
	var buf bytes.Buffer

	w := zip.NewWriter(&buf)

	for _, fc := range content {
		f, err := w.Create(fc.Name)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(fc.Body)
		if err != nil {
			return nil, err
		}
	}

	w.Close()

	return bytes.NewReader(buf.Bytes()), nil
}
